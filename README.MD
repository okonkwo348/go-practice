# ASCII Art Web

Ascii-art -web consists in creating and running a server in which it use a web GUI (graphical user interface) to display a graphical representation of a text given in a selected banner style.

## Description
Ascii-art -web consists in creating and running a server in which it use a web GUI (graphical user interface) to display a graphical representation of a text given in a selected banner style. The text can be letters (a, b, c, d,..), numbers (0, 1, 2, 3 ..9), punction ( , . ? !), and symbols (~ @ # $ 5 ^ &) printed in either three banner styles; standard, shadow and thinkertoy. Each character has eight rows and a new line.

- Key Features:

***Logic***
validateInput function
LoadBanner Function
getCharLine function
renderWord function
renderAll function

***Template***
home.html
result.thml

***Handler**
homeHandler function
asciiArtHandler function

***Path***
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

***Execution Command**
to get the server started:  go run .

***Example***
Open this http://localhost:8080  on browser.
Write a text at the input text
Select either Standard, Shadow or Thinkertoy.
Click Submit butten
Click Reset Button to reset text and banner style



## Implementation 

***Logic Overview***
- 1. Type a text in the input text and select a banner style on a web home page form.
- 2. The text and banner style selected is received as a request by the HTTP GET method through the current path and it called at homeHandler function after the function template.ParseFile is used to locate the home.html file.
- 3. The content of the file is place to a variable w which is later executed and sent as a http response.
- 4. 

