// Package jwt is a jwt implementation of the auth interface
package jwt

import (
	"sync"
	"time"

	"github.com/go-iot-platform/go-micro/auth"
	"github.com/go-iot-platform/micro-plugins/auth/jwt/token"
	"github.com/go-iot-platform/micro-plugins/auth/jwt/token/jwt"
)

// NewAuth returns a new instance of the Auth service
func NewAuth(opts ...auth.Option) auth.Auth {
	j := new(jwtAuth)
	j.Init(opts...)
	return j
}

func NewRules() auth.Rules {
	return &jwtRules{}
}

type jwtAuth struct {
	sync.Mutex
	options auth.Options
	token   token.Provider
}

type jwtRules struct {
	sync.Mutex
	rules []*auth.Rule
}

func (j *jwtAuth) String() string {
	return "jwt"
}

func (j *jwtAuth) Init(opts ...auth.Option) {
	j.Lock()
	defer j.Unlock()

	for _, o := range opts {
		o(&j.options)
	}

	j.token = jwt.NewTokenProvider(
		token.WithPrivateKey(j.options.PrivateKey),
		token.WithPublicKey(j.options.PublicKey),
	)
}

func (j *jwtAuth) Options() auth.Options {
	j.Lock()
	defer j.Unlock()
	return j.options
}

func (j *jwtAuth) Generate(id string, opts ...auth.GenerateOption) (*auth.Account, error) {
	options := auth.NewGenerateOptions(opts...)
	if len(options.Issuer) == 0 {
		options.Issuer = j.Options().Issuer
	}
	name := options.Name
	if name == "" {
		name = id
	}
	account := &auth.Account{
		ID:       id,
		Type:     options.Type,
		Scopes:   options.Scopes,
		Metadata: options.Metadata,
		Issuer:   options.Issuer,
		Name:     name,
	}

	// generate a JWT secret which can be provided to the Token() method
	// and exchanged for an access token
	secret, err := j.token.Generate(account, token.WithExpiry(time.Hour*24*365))
	if err != nil {
		return nil, err
	}
	account.Secret = secret.Token

	// return the account
	return account, nil
}

func (j *jwtAuth) Inspect(token string) (*auth.Account, error) {
	return j.token.Inspect(token)
}

func (j *jwtAuth) Token(opts ...auth.TokenOption) (*auth.Token, error) {
	options := auth.NewTokenOptions(opts...)

	secret := options.RefreshToken
	if len(options.Secret) > 0 {
		secret = options.Secret
	}

	account, err := j.token.Inspect(secret)
	if err != nil {
		return nil, err
	}

	access, err := j.token.Generate(account, token.WithExpiry(options.Expiry))
	if err != nil {
		return nil, err
	}

	refresh, err := j.token.Generate(account, token.WithExpiry(options.Expiry+time.Hour))
	if err != nil {
		return nil, err
	}

	return &auth.Token{
		Created:      access.Created,
		Expiry:       access.Expiry,
		AccessToken:  access.Token,
		RefreshToken: refresh.Token,
	}, nil
}

func (j *jwtRules) Grant(rule *auth.Rule) error {
	j.Lock()
	defer j.Unlock()
	j.rules = append(j.rules, rule)
	return nil
}

func (j *jwtRules) Revoke(rule *auth.Rule) error {
	j.Lock()
	defer j.Unlock()

	rules := []*auth.Rule{}
	for _, r := range j.rules {
		if r.ID != rule.ID {
			rules = append(rules, r)
		}
	}

	j.rules = rules
	return nil
}

func (j *jwtRules) Verify(acc *auth.Account, res *auth.Resource, opts ...auth.VerifyOption) error {
	j.Lock()
	defer j.Unlock()

	var options auth.VerifyOptions
	for _, o := range opts {
		o(&options)
	}

	return auth.VerifyAccess(j.rules, acc, res)
}

func (j *jwtRules) List(opts ...auth.RulesOption) ([]*auth.Rule, error) {
	j.Lock()
	defer j.Unlock()
	return j.rules, nil
}
