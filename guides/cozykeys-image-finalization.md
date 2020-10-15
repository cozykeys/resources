# How to Finalize Images for the CozyKeys Website

```bash
# Install ImageMagick; on Ubuntu:
sudo apt install -y imagemagick

# At the time of writing, both of my previous phones take pictures at 3024x3024
# but this may not always be true. Check the source image sizes via:
identify -format "%wx%h" image.jpg

# Make sure the image is saved as PNG, resize it to the desired sizes, and fix
# any orientation quirks:
convert image.jpg -resize 3024x3024 -auto-orient image_3024x3024.png
convert image.jpg -resize 1600x1600 -auto-orient image_1600x1600.png
convert image.jpg -resize 800x800 -auto-orient image_800x800.png

# Upload to S3 and allow public read

# Upload to Google Drive as well
```
