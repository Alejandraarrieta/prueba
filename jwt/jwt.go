package jwt
import(
	"crypto/rsa"
	"io/ioutil"
	"time"
	"log"
	jwt "github.com/golang-jwt/jwt/v4"
)

type ResponseToken struct{
	Token string `json:"token"`
}

var privateKey *rsa.PrivateKey
var publicKey *rsa.PublicKey

//func ParseRSAPrivateKeyFromPEM(key []byte) (*rsa.PrivateKey, error)
func init(){
	privateBytes, err := ioutil.ReadFile("./private.rsa") //crear key y pasar ruta del archivo
	if err != nil{
		log.Fatal("could not read the file:",err)
	}
	publicBytes, err := ioutil.ReadFile("./public.rsa.pub") //crear key y pasar ruta del archivo
	if err != nil{
		log.Fatal("could not read the file", err)
	}

	privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateBytes)
	if err != nil{
		log.Fatal("Could not parse privatekey")
	}
	publicKey, err = jwt.ParseRSAPublicKeyFromPEM(publicBytes)
	if err != nil{
		log.Fatal("Could not parse publickey")
	}
}

func GenerateJWT(user UserData) (string, error){
	token:=jwt.New(jwt.GetSigningMethod("RS256"))
	token.Claims = jwt.MapClaims{
		"exp":  time.Now().Add(time.Hour * 1).Unix(),
		"iat":  time.Now().Unix(),
		"name": user.Name,
		"lastname": user.Lastname,
		"email": user.Email,
		"phone": user.Phone,
		"area": user.Area,
		"workunit": user.Workunit,
	}
	result, err := token.SignedString(privateKey)
	return result, err
}