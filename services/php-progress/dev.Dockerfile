# Use the official PHP image with FPM
FROM php:8.2-fpm

# Install necessary system packages and PHP extensions
RUN apt-get update && apt-get install -y \
    libssl-dev \
    libzip-dev \
    unzip \
    git \
    && pecl install mongodb \
    && docker-php-ext-enable mongodb \
    && docker-php-ext-install zip

# Set the working directory
WORKDIR /var/www/html

COPY . .

# Expose port 9000 for FPM
EXPOSE 9000

# By default, start php-fpm
CMD ["php-fpm"]
