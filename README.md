# Image to ASCII

This Go project converts images into ASCII art representations.

## Features

- Supports various image formats (PNG, JPEG, GIF)
- Allows resizing the image before conversion
- Offers different interpolation methods for resizing
- Provides an option to reverse the tonal values of the resulting ASCII art

## Usage

```Bash
image-to-ascii [OPTIONS] [PATH] [RESIZE_FACTOR]
```

## Arguments

- `PATH`: Path to the image file you want to convert.
- `RESIZE_FACTOR`: A floating-point number representing the resize factor. Defaults to 1 (no resize).

### Options

- `-r` or `--reverse`: Reverses the tonal values of the resulting ASCII art (dark becomes light, and vice versa).
- `--l2`: Uses Lanczos2 interpolation for resizing.
- `--l3`: Uses Lanczos3 interpolation for resizing.
- `--bc`: Uses Bicubic interpolation for resizing (default).
- `--bl`: Uses Bilinear interpolation for resizing.
- `--nn`: Uses Nearest Neighbor interpolation for resizing.
- `--mn`: Uses Mitchell Netravali interpolation for resizing.
- `-h` or `--help`: Displays usage information.

### Example

```Bash
image-to-ascii -r photo.jpg 0.5  # Convert photo.jpg to ASCII art with resize factor 0.5 and reversed tones.
```

## Dependencies

This project requires Go 1.22 and uses the following external library:
- `github.com/nfnt/resize` for image resizing functionality.
