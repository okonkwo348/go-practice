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
- 4. The two inputs (text and banner style) from the home page is gotten is sent as a request to the asciiArthandler function. 
- 5. The text is check to ensure it is not empty and later the it splited so it can be validated as an ascii character. 
- 6. Check if the banner style is empty. If the banner style is not empty then :oad the banner file. 
- 7. The validate ascii text is taken and each of it equavalent ascii art characters will be gotten a stored in a variable. 
- 8. The result.hmtl file is parsed and found.
- 9. The generated art which is stored in a variable is called in the result.html file in a dynamic action and it is displayed by the url path linked to the asciiArthandler function.



The link between the two file is found at the form beginng tag at the attribute action which send the two inputs (textinput and banner style) to the url path of asciiArtHandler. This automaticall connects the two handler functions because the inputs from response from homeHandler is sent to "/ascii-art" url path and asciiArtHandle function receives it as it request.




