El documento esta en:
https://docs.google.com/document/d/1iUem-Yt4eihj-WmNQ-xER8gspANHc9YDrSWZyWB2H94/edit
<br />
https://github.com/GoesToEleven/golang-arch

# HMAC
 prevent falso bearer token, usando cryptographic "signing"<br />
 es un hash basado en el mensaje, codigo autentificacion y algunas otras cosas

# AES
Encripta y desecripta cada vez que ejecutamos la funcion, no necesito un desecriptador y un encriptador<br />
Una vez lo encrtipta, la segunda lo desecripta, tercera lo encripta y asi susesivamente

 # JWT
```javascript
{JWT standard fields}.{Your fields}.Signature
```
```sh
github.com/dgrijalva/jwt-go
```

- MD5 (NO Usar, tiene vulnerabilidades)
- SHA
- BCrypto
- Scrypto
- Rot13: Encripta el mensaje cambiando de lugar los caracteres
<br />

![Events](../images/61.png)


# OAuth2
Autentificacion en otro sistema, no en el nuestro <br />
Autoriza a un usuario a hacer cosas en otro sitio<br />
Lo mas comun es para el login <br />
Tiene 3 cosas:
- Credenciles del cliente
    - mas seguro, pero requiere del server
- Implicit
    - menos seguro

<br />
golang.org/x/oauth2 <br/>
Es muy bueno pero no tanto para los que no tiene <br/>
https://godoc.org/golang.org/x/oauth2#pkg-subdirectories <br />

- Va a tu sitio
- luego a google
- google redirecciona a tu sitio con un query string con el codigo que necesita el sitio
<br />
- user is as spacex.com (for example)
    - logs in with Oauth2 using google Oauth2
    - Redirects user to Google Oauth login page
        - user is asked to grant permissions
        - what to share from google account
    - google Redirects back to spacex.com with a code
    - Spacex.com exchanges code and secret for access token to google
    - Spacex.com uses token to get who the user is on google, including user id on google
<br />
Por ejemplo para github vamos a OAuth Apps > Register a new application<br />
Home page: el home de la pagina <br />
Authorization callback URL: donde recibiras el redirect con el codigo <br />
Obtenemos el Client ID y el Client Secret