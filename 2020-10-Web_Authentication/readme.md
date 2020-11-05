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