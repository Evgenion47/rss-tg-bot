# homework-2

Телеграм бот который присылает обновления из RSS фидов на которые подписался пользователь.

Схема архитектуры приложения:
![](/My_First_Board.jpg "схема архитектуры приложения")

Управление происходит через http/gRPC ручки и бота telegram который является человеко читаемым интерфейсом. При добавлении нового источника в пул сервис UI шлет сигнал сервису RSS сборщику чтобы он обновил список источников. При получении новости RSS сборщик посылает сигнал сервису UI о получении новой новости чтобы тот проверил кому ее необходимо отправить. 

Схема базы данных:
![](/DB.png "схема базы данных")

Файл миграции для базы данных находится в репозитории.
