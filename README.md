Не успел/можно было бы сделать лучше:


1) И dto, и доменные сущности хранятся в models. Не самое лучшее решение, 
но придерживаться слоистои архитектуры было бы избыточно

2) Логировать с помошью обертки над logger (например slog), 
например чтобы контроллировать куда логировать в зависимости от окружения(local, dev, prod)

3) Добавить фабричные методы для структур e.g fromXmlCurrency()

4) Проблема с кодировкой mssql ???

5) Не хватило времени подумать какой контекст передать в repository, gorilla.context очевидно не подходит
так как закрывается после отдачи http ответа

6) не успел разобраться почему документация не сгенерировалась, с chi работало без проблем 
