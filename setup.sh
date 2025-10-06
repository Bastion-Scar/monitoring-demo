#!/bin/bash
# setup.sh - Интерактивная настройка проекта

# Функция проверки ошибки
check_error() {
    local cmd="$@"
    eval "$cmd"
    if [ $? != 0 ]; then
        echo "Ошибка при выполнении команды '$cmd'." >&2
        exit 1
    fi
}

# Регулярные выражения для валидации
REGEX_TELEGRAM_TOKEN='^[0-9]{10}:.{35}$'
REGEX_TELEGRAM_CHAT_ID='^-?[0-9]+$'

# Проверяем наличие файла конфигурации
if [ -f "config/tg.env" ]; then
    echo "🚧 ⚠️ Файл 'tg.env' уже существует!"
    read -rp "Перезаписать конфигурационный файл? (y/N): " overwrite_tg
    if [[ ! "${overwrite_tg}" =~ ^[Yy]$ ]]; then
        echo "Настройка файла tg.env отменена."
    fi
else
    # Сбор необходимых данных от пользователя для Telegram
    echo "🛠 🚀 Настройка мониторинг-демо"
    echo "=========================================="

    echo "📞 Введите необходимые данные для Telegram бота:"

    # Получение и проверка бот-токена
    while true; do
        read -rp "Telegram Bot Token: " bot_token
        if [[ -z "$bot_token" ]]; then
            echo "❗ Бот-токен не может быть пустым!"
        elif ! [[ "$bot_token" =~ $REGEX_TELEGRAM_TOKEN ]]; then
            echo "❗ Неправильный формат токена! Токен должен содержать примерно 'число:строка'. Попробуйте ещё раз."
        else
            break
        fi
    done

    # Получение и проверка чат-ID
    while true; do
        read -rp "Telegram Chat ID: " chat_id
        if [[ -z "$chat_id" ]]; then
            echo "❗ Чат-ID не может быть пустым!"
        elif ! [[ "$chat_id" =~ $REGEX_TELEGRAM_CHAT_ID ]]; then
            echo "❗ Неверный формат Chat ID! Должен быть целым числом (например, '-12345'). Попробуйте ещё раз."
        else
            break
        fi
    done

    # Создание файла tg.env
    cat > config/tg.env << EOF
# Telegram Bot Configuration
TELEGRAM_BOT_TOKEN="${bot_token}"
TELEGRAM_CHAT_ID="${chat_id}"
EOF

    check_error "cat config/tg.env"
    echo "✅ Файл 'tg.env' создан."
fi


# Проверяем наличие файла db.env
if [ -f "config/db.env" ]; then
    echo "🚧 ⚠️ Файл 'db.env' уже существует!"
    read -rp "Перезаписать конфигурационный файл? (y/N): " overwrite_db
    if [[ ! "${overwrite_db}" =~ ^[Yy]$ ]]; then
        echo "Настройка файла db.env отменена."
    fi
else
    # Сбор необходимых данных для базы данных
    echo "🔍 Введите необходимые данные для базы данных MySQL:"

    # Получение корня пароля
    while true; do
        read -rp "MySQL Root Password: " MYSQL_ROOT_PASSWORD
        if [[ -z "$MYSQL_ROOT_PASSWORD" ]]; then
            echo "❗ Пароль не может быть пустым!"
        else
            break
        fi
    done

    # Название базы данных
    while true; do
        read -rp "Название базы данных (database): " MYSQL_DATABASE
        if [[ -z "$MYSQL_DATABASE" ]]; then
            echo "❗ Название базы данных не может быть пустым!"
        else
            break
        fi
    done

    # Пользователь MySQL
    while true; do
        read -rp "Пользователь MySQL: " MYSQL_USER
        if [[ -z "$MYSQL_USER" ]]; then
            echo "❗ Имя пользователя не может быть пустым!"
        else
            break
        fi
    done

    # Пароль пользователя MySQL
    while true; do
        read -rp "Пароль пользователя MySQL: " MYSQL_PASSWORD
        if [[ -z "$MYSQL_PASSWORD" ]]; then
            echo "❗ Пароль не может быть пустым!"
        else
            break
        fi
    done

    # Создание файла db.env
    cat > config/db.env << EOF
# MySQL Configuration
MYSQL_ROOT_PASSWORD="${MYSQL_ROOT_PASSWORD}"
MYSQL_DATABASE="${MYSQL_DATABASE}"
MYSQL_USER="${MYSQL_USER}"
MYSQL_PASSWORD="${MYSQL_PASSWORD}"
EOF

    check_error "cat config/db.env"
    echo "✅ Файл 'db.env' создан."

    # Генерация строки подключения (DSN) для app.env
    DSN="${MYSQL_USER}:${MYSQL_PASSWORD}@tcp(db:3306)/${MYSQL_DATABASE}?charset=utf8mb4&parseTime=True&loc=Local"

    # Создание файла app.env
    cat > config/app.env << EOF
# App Configuration
DSN="${DSN}"
EOF

    check_error "cat config/app.env"
    echo "✅ Файл 'app.env' создан."
fi

# Поднимаем контейнеры Docker Compose
echo "🚀 Запускаем проект..."
check_error "docker compose up -d"

exit 0