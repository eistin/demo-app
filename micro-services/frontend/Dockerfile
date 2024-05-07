# Use the official PHP Apache image
FROM php:apache

# Set the working directory
WORKDIR /var/www/html

# Copy the PHP files into the container
COPY . .

# Expose port 80
EXPOSE 80

# Start Apache server
CMD ["apache2-foreground"]