package service

import (
	"context"
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/yanarowana123/kmf/internal/model"
	"io"
	"log"
	"net/http"
	"time"
)

type CurrencyRepository interface {
	Save(ctx context.Context, currency model.SaveCurrency) error
	List(ctx context.Context, date time.Time, code string) ([]model.GetCurrencyResponse, error)
}

type CurrencyService struct {
	r   CurrencyRepository
	url string
}

func NewCurrencyService(r CurrencyRepository, url string) CurrencyService {
	return CurrencyService{r, url}
}

func (s CurrencyService) Save(ctx context.Context, date time.Time) error {
	xmlBytes, err := s.request(ctx, date)
	if err != nil {
		log.Println(err.Error())
		return errors.New("failed to save currency data")
	}

	var currencyXml model.CurrencyXml
	err = xml.Unmarshal(xmlBytes, &currencyXml)
	if err != nil {
		log.Println(err.Error())
		return errors.New("failed to save currency data")
	}

	go s.save(ctx, currencyXml)

	return nil
}

func (s CurrencyService) request(ctx context.Context, date time.Time) ([]byte, error) {
	response, err := http.Get(fmt.Sprintf("%s?fdate=%s", s.url, date.Format("02.01.2006")))
	if err != nil {
		return []byte{}, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return []byte{}, fmt.Errorf("unexpected status code %d", response.StatusCode)
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return []byte{}, fmt.Errorf("Read body: %v", err)
	}

	return data, nil
}

func (s CurrencyService) save(ctx context.Context, currencyXml model.CurrencyXml) {
	for _, item := range currencyXml.Currencies {
		date, err := time.Parse("02.01.2006", currencyXml.Date)
		if err != nil {
			log.Println(err.Error())
			return
		}
		saveCurrency := model.SaveCurrency{
			Date:  date,
			Title: item.Title,
			Code:  item.Code,
			Value: item.Value,
		}
		go func() {
			err = s.r.Save(ctx, saveCurrency)
			if err != nil {
				log.Println(err.Error())
			}
		}()
	}
}

func (s CurrencyService) List(ctx context.Context, date time.Time, code string) ([]model.GetCurrencyResponse, error) {
	currencies, err := s.r.List(ctx, date, code)
	if err != nil {
		log.Println(err.Error())
	}
	return currencies, err
}
