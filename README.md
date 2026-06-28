# ASCII Art Web

Ascii-art-web consists in creating and running a server in which it use a web GUI (graphical user interface) to display a graphical representation of a given text in a selected banner style.

## Description
Ascii-art -web consists in creating and running a server in which it use a web GUI (graphical user interface) to display a graphical representation of a given text in a selected banner style. The text can be letters (a, b, c, d,..), numbers (0, 1, 2, 3 ..9), punction ( , . ? !), and symbols (~ @ # $ * ^ &) printed in either three banner styles; standard, shadow and thinkertoy. Each character has eight rows and a new line.

- Key Features:

***Logic***
validateInput function
LoadBanner Function
getCharLine function
renderWord function
renderAll function

***Template***
home.html
result.html

***Handler***
homeHandler function
asciiArtHandler function

***url Path***
http.HandleFunc("/", homeHandler)
http.HandleFunc("/ascii-art", asciiArtHandler)

***Technologies***
HTML
GOLANG



## Authors

***Contributor***
Okonkwo Paul

***Contact***
okonkwopaul348@gmail.com

## Usage

***Execution Command***
To get the server started:  go run .

***Example***
- Open this http://localhost:8080  on browser.
- Write a text at the input text
- Select either Standard, Shadow or Thinkertoy.
- Click Submit button.
- Click Reset button to reset text and banner style



## Implementation Details

***Logic Overview***
- 1. Type a text in the input text and select a banner style on a web home page form at http://localhost:8080 browser.
- 2. The homeHandler function uses template.ParseFiles func to locate the home.html file.
- 3. The content of the file is place to in a variable w which is later executed and sent as a http response.
- 4. The two inputs (text and banner style) from the home page is gotten through "/ascii-art" action at the form tag and sent as a request to the asciiArthandler function. 
- 5. The asciiArthandler function check the text if it is not empty afterward split it so can be validated as an ascii character. 
- 6. Check if the banner style is empty. If the banner style is not empty then load the banner file. 
- 7. The validate ascii text is taken and each of it equavalent ascii art characters will be gotten and stored in a variable called generate. 
- 8. The result.hmtl file is parsed and found to be able to generate the ascii art.
- 9. The generated art which is stored in a variable is called in the result.html file in a dynamic action and it is displayed at the url path linked to the asciiArthandler function.

***Important note***
The action="/ascii-art" in the home.html form is the bridge between everything. That one attribute:

- Links home.html to result.html
- Links homeHandler to asciiArtHandler
- Links the frontend to the backend




