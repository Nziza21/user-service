module github.com/Nziza21/user-service

go 1.18

replace github.com/chenzhuoyu/iasm => github.com/chenzhuoyu/iasm v0.8.0

require (
	github.com/gin-gonic/gin v1.9.0
	github.com/google/uuid v1.3.0
	github.com/joho/godotenv v1.5.1
	github.com/swaggo/files v0.0.0-20220610200504-28940afbdbfe
	github.com/swaggo/gin-swagger v1.5.0
	github.com/swaggo/swag v1.8.1
	golang.org/x/crypto v0.13.0
	gorm.io/driver/postgres v1.5.0
	gorm.io/gorm v1.25.1
)

require github.com/golang-jwt/jwt/v4 v4.5.2 // indirect

require (
	github.com/KyleBanks/depth v1.2.1
	github.com/PuerkitoBio/purell v1.1.1
	github.com/PuerkitoBio/urlesc v0.0.0-20170810143723-de5bf2ad4578
	github.com/bytedance/sonic v1.8.0
	github.com/chenzhuoyu/base64x v0.0.0-20221115062448-fe3a3abad311
	github.com/gin-contrib/sse v0.1.0
	github.com/go-openapi/jsonpointer v0.19.5
	github.com/go-openapi/jsonreference v0.19.6
	github.com/go-openapi/spec v0.20.4 
	github.com/go-openapi/swag v0.19.15 
	github.com/go-playground/locales v0.14.1 
	github.com/go-playground/universal-translator v0.18.1 
	github.com/go-playground/validator/v10 v10.11.2 
	github.com/goccy/go-json v0.10.0 
	github.com/jackc/pgpassfile v1.0.0 
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a 
	github.com/jackc/pgx/v5 v5.3.0 
	github.com/jinzhu/inflection v1.0.0 
	github.com/jinzhu/now v1.1.5 
	github.com/josharian/intern v1.0.0 
	github.com/json-iterator/go v1.1.12 
	github.com/klauspost/cpuid/v2 v2.0.9 
	github.com/leodido/go-urn v1.2.1 
	github.com/mailru/easyjson v0.7.6 
	github.com/mattn/go-isatty v0.0.17
	github.com/modern-go/concurrent v0.0.0-20180228061459-e0a39a4cb421 
	github.com/modern-go/reflect2 v1.0.2 
	github.com/pelletier/go-toml/v2 v2.0.6 
	github.com/twitchyliquid64/golang-asm v0.15.1 
	github.com/ugorji/go/codec v1.2.9 
	golang.org/x/arch v0.0.0-20210923205945-b76863e36670 
	golang.org/x/net v0.10.0 
	golang.org/x/sys v0.12.0 
	golang.org/x/text v0.13.0 
	golang.org/x/tools v0.6.0 
	google.golang.org/protobuf v1.28.1 
	gopkg.in/yaml.v2 v2.4.0 
	gopkg.in/yaml.v3 v3.0.1 
)
