
# ASCII Art Web Application

The ASCII Art Web Application allows users to generate ASCII art from text input. Users can enter text, select a font style, and the application will convert the text into an ASCII art representation. The generated ASCII art can be displayed on the web page and downloaded as a text file.

## Features
- User-friendly web interface
- Text input field with a maximum length of 500 characters
- Three font style options: Standard, Shadow, and Thinkertoy
- ASCII art generation based on user input and selected font style
- Display of generated ASCII art on the web page
- Download functionality to save the ASCII art as a text file

## Installation
1. Clone the repository:
   ```
   git clone https://learn.reboot01.com/git/salkoohe/ascii-art-stylize.git
   ```

2. Navigate to the project directory:
   ```
   cd ASCII-ART-STYLIZE
   ```

3. Build and run the application:
   ```
   go run . 
   
   ```

4. The application will start running on [http://localhost:8080](http://localhost:8080).

## Usage
1. Open the application in your web browser by visiting [http://localhost:8080](http://localhost:8080).
2. Enter the desired text in the textarea field.
3. Select the font style you prefer by clicking one of the radio buttons (Standard, Shadow, or Thinkertoy).
4. Click the "Draw" button to generate the ASCII art.
5. The generated ASCII art will be displayed on the web page.
6. To download the ASCII art as a text file, click the download button (the button with the download icon).

## Project Structure
- `main.go`: The entry point of the application, which starts the server.
- `functions/server.go`: Contains the server setup and handler functions for different routes.
- `static/index.html`: The HTML template for the web interface.
- `static/style.css`: The CSS file for styling the web interface (not included in the provided code).
- `static/404.html`, `static/400.html`, `static/500.html`: Error pages for different HTTP status codes (not included in the provided code).
- `static/download.png`: The image used for the download button (not included in the provided code).

## Dependencies
- Go programming language

## Contributing
Contributions are welcome! If you find any issues or have suggestions for improvements, please open an issue or submit a pull request.

## License
Open source