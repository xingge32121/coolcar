/*
 * @Author: Wujiahuo
 * @Date: 2022-09-02 11:33:32
 * @LastEditors: OBKoro1
 * @LastEditTime: 2022-09-05 11:44:16
 * @FilePath: \wxc:\Users\Administrator\go\src\coolcar\server\shared\token\auth\verfily_test.go
 * @Description:
 * Copyright (c) 2022 by Wujiahuo/bzlrobot, All Rights Reserved.
 */
package auth

import (
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const publicKey = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA2wDr1w5df1HnC8TzQyy/
5XPbAmqb/+8/S9L6D6CBDRT7HY2JeJw6ZSeYdYaUX0vJtngppMmOJFSanVAVCJHZ
laFNMlgf+P1ebX8FTOBPIOX1aANss6PuFj/S3+MDJ9CtomQ3h47gRrdiYJStlRn5
j1POwjU2O3ZxkoqQuwqt9/S3duqXTZEnGcSc/oUlMuw217u82OSgZzZuiNxLSlQ0
15BwnLJOt6ZWETWe9P+tRNKsxaGaoLK/5PDF8mK3cXTr0YVYJeNppRUJqt/LQGzl
yBB3khTIAa0LVHEloZOUDQW7vl/RMbMvtq+ZG7xBYLUPqdRBjq9+evY4nbydyD7f
tQIDAQAB
-----END PUBLIC KEY-----`

func TestVerfily(t *testing.T) {
	// pubKeys, err := os.Open("../../../auth/public.key")
	// if err != nil {
	// 	t.Fatalf("cannot Open pubKeys %v", err)
	// }
	// pubBytes, err := ioutil.ReadAll(pubKeys)
	// if err != nil {
	// 	t.Fatalf("cannot ReadAll pubBytes %v", err)
	// }
	pubKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(publicKey))
	if err != nil {
		t.Fatalf("cannot Parse pubKey %v", err)
	}
	tk := &JWTTokenGenVerfily{
		PublicKey: pubKey,
	}
	cases := []struct {
		name    string
		tkn     string
		now     time.Time
		want    string
		wantErr bool
	}{
		{
			name: "valid_token",
			tkn:  `eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjIxMDg5NzksImlhdCI6MTY2MjEwNTg5NSwiaXNzIjoiY29vbGNhci9hdXRoIiwic3ViIjoiNjMwODg5YTkzMzk2ZmEwMjMyZDVmM2Q4In0.0m0_w8JQrRa_eizRSRYmwBNhyDWBerxnedGG61VxIsxwsSMUtUHxGPKpuVOHe37x6Cxxlkv9gOhX8krr4IpjYuXP2UX60tzepuDXcI3XUigFAZcBoVVC7r1DwwkGs5oicJIiq-q5ZcIDi012gu1ot78TPpEap7pHXEfp6NQQoThQaFBb0_-0C5nBwEe0ikUcwHVQCcXxqISsdCuCvRY2twtPbPBsMwV6C2aOWPlwlgNQ_ZrheeYwrtx6tI_hmzAW4WRuANbeFjR4PbLDODpfwNXPS5sO4YX9bZN30Ii7XLdiIebo9NGQ-2A8q9niIgL4TeDWBxF_WG4pUfBkNZAD0g`,
			now:  time.Unix(1662105895, 0),
			want: "630889a93396fa0232d5f3d8",
		},
		{
			name:    "valid_expired",
			tkn:     `eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjIxMDM5NzksImlhdCI6MTY2MjEwMDM3OSwiaXNzIjoiY29vbGNhci9hdXRoIiwic3ViIjoiNjMwODg5YTkzMzk2ZmEwMjMyZDVmM2Q4In0.B2ZvuY26XgVjTqRhbz2Lu6mop9yzY7S6erAdd1BHJoxl7jrr285P91jHh-VOdzw_cEeXoJN64ccLGw3siaAx1HqDt_RCsC53GgqYfPv8lQvQLP1iJkSEm_xi6AlBBRSrMnIePtPkbkj6taMlkEYt5bYAu5h9vy0qK74_9yC1RZ1qlqyTPqTEsfvI-93s4Z2kCBWLMiSGjK3pRwmtHzHfn0J_3z7OdMjQCeOlfbrd4v3da35ieX8_RQTPalnCqj8ISMi1JvC8YT3T45I82nZcPeTWAy8I9AoYZagbSUBkSIcQfpkJVrdXssOXWQNF0Q_Ia5slGQD1jN7rJIs9Pg99Yg`,
			now:     time.Unix(1660100379, 0),
			wantErr: true,
		}, {
			name:    "valid_bad",
			tkn:     `valid_bad.eyJleHAiOjE2NjIxMDM5NzksImlhdCI6MTY2MjEwMDM3OSwiaXNzIjoiY29vbGNhci9hdXRoIiwic3ViIjoiNjMwODg5YTkzMzk2ZmEwMjMyZDVmM2Q4In0.B2ZvuY26XgVjTqRhbz2Lu6mop9yzY7S6erAdd1BHJoxl7jrr285P91jHh-VOdzw_cEeXoJN64ccLGw3siaAx1HqDt_RCsC53GgqYfPv8lQvQLP1iJkSEm_xi6AlBBRSrMnIePtPkbkj6taMlkEYt5bYAu5h9vy0qK74_9yC1RZ1qlqyTPqTEsfvI-93s4Z2kCBWLMiSGjK3pRwmtHzHfn0J_3z7OdMjQCeOlfbrd4v3da35ieX8_RQTPalnCqj8ISMi1JvC8YT3T45I82nZcPeTWAy8I9AoYZagbSUBkSIcQfpkJVrdXssOXWQNF0Q_Ia5slGQD1jN7rJIs9Pg99Yg`,
			now:     time.Unix(1660100379, 0),
			wantErr: true,
		}, {
			name:    "valid_err",
			tkn:     `eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjIxMDM5NzksImlhdCI6MTY2MjEwMDM3OSwiaXNzIjoiY29vbGNhci9hdXRoIiwic3ViIjoiNjMwODg5YTkzMzk2ZmEwMjMyZDVmM2Q2In0.B2ZvuY26XgVjTqRhbz2Lu6mop9yzY7S6erAdd1BHJoxl7jrr285P91jHh-VOdzw_cEeXoJN64ccLGw3siaAx1HqDt_RCsC53GgqYfPv8lQvQLP1iJkSEm_xi6AlBBRSrMnIePtPkbkj6taMlkEYt5bYAu5h9vy0qK74_9yC1RZ1qlqyTPqTEsfvI-93s4Z2kCBWLMiSGjK3pRwmtHzHfn0J_3z7OdMjQCeOlfbrd4v3da35ieX8_RQTPalnCqj8ISMi1JvC8YT3T45I82nZcPeTWAy8I9AoYZagbSUBkSIcQfpkJVrdXssOXWQNF0Q_Ia5slGQD1jN7rJIs9Pg99Yg`,
			now:     time.Unix(1662100379, 0),
			wantErr: true,
		},
	}
	for _, v := range cases {
		t.Run(v.name, func(t *testing.T) {
			jwt.TimeFunc = func() time.Time {
				return v.now
			}
			accounId, err := tk.Verfily(v.tkn)
			if !v.wantErr && err != nil {
				t.Errorf("verification failed: %v", err)
			}

			if v.wantErr && err == nil {
				t.Errorf("want error; got no error")
			}
			if accounId != v.want {
				t.Errorf("wrong want verfily. want: %q; got: %q", v.want, accounId)
			}
		})
	}
}
