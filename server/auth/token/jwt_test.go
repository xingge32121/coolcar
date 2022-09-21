/*
 * @Author: Wujiahuo
 * @Date: 2022-09-01 10:52:38
 * @LastEditors: OBKoro1
 * @LastEditTime: 2022-09-14 17:49:36
 * @FilePath: \wxc:\Users\Administrator\go\src\coolcar\server\auth\token\jwt_test.go
 * @Description:
 * Copyright (c) 2022 by Wujiahuo/bzlrobot, All Rights Reserved.
 */
package token

import (
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
)

/**
 * @description:
 * @return {*}
 */
const privateKey = `-----BEGIN PRIVATE KEY-----
MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQCqvwbo4Zoo0cxA
m6UtFjr0sA7BdyKWOg1C0iOL083eyChrxZXC+bWW87BXBEWeNXi5yuAQ4yPzkNaF
8vq87BQiW5bPRiZyfj8q5fg8e6/9Is8tiKtIBu55qZgdajHRD9X3J3FKkzGIVFbi
2C3hZGIGhst+n8x6HWAHp1HkqTD+f5zewUft27s79YsSPq8Ub3pDdWdBqyXNvwxK
OBRTIPmZgxx3amiYC8CVeRRpnMdsM/Kafmq1Pq30fBXLOezeM0I0cjQJJiMewLln
8JoXROr+pBne5V3vR6bUHrdSuHuzCOvDna5oQX3tZHvzS/zofVXusESCcIzrBH5g
lR6hOWmpAgMBAAECggEAYtVmO4DzCfO+phsx9jIq+B28vNGoDIsXx/j5KGP2g6Ea
fJ+HO6/tI651ATlwzs4mzmyXXE8OLVtGQYzPBHImiWi4l8wETzuydrNdCC2URhNT
FwlIE84nVfcMHOKuaVQgRwrmsgMXpdEj5fS6QgAoGZ0Gx4naA0ljn6qDZbdlq7nh
+akc0HQ2kOYt0AAt6Hv2xaoZ+eq4/M7OMM6nyIoSWBcFsL1t4S5ogqQ3BSuVYK38
hX/5/daZyQR4HU/suPDNb04k+GSFBC5lGTRp7KEJ9Z08Q/p3RPH9AFJc9jH4rVge
BdSocqrPNZB3qKo3F2XD5X67MAeUnE5alxXJOTTC7QKBgQDdGlY624mMKPiGQa6P
GnCfBeCPh2LwCHy6bSW0waLQlnwWMV9vxLi/S+/r6yVtEUmiiSiOAKLkFNKNubVo
YS6+M5iZnD00ve9ZpAFnJlkc12jF84wjNd4cFfl6AQWx1izrTJUw/ydIngrxy3bQ
uEKRfGO54W062MPM7z8d4wRHBwKBgQDFsgc5yg8jd1hmkFhsojvQFoXKGkC9f0sP
F0VYvNazJnJ+eH/uSYG7oUzlusk1COgM+o5lk9dVBW5k/6CeGXERUFQeSQaBrxir
PDpGHqX++uuoPSyufy3AatWNfHae664jJush7WAwmNk+nroDEraND4uuel3VFdDr
mqWjdcptzwKBgQDV3cPK6tABYyABvUa75TdKmsS8EJkC1TLQoQL39NUoLmeYj+lg
k0igwTMxYCKzfSVcso8nzDXDqOCBdkCi1l4AXKge7aHGgnIioyE4RT+tRi9ySkIj
TNerExZdjN+VRTXt5AXrfuLguv6bI9Op0JGTaF6OTh757PJ4KHNE1XItwQKBgB0V
5i5AZh3Hrz/XuMqyy63/wktLtX6rvbVIrTEzBztwSqQEwn2iXu3l+1RJONUOrGVM
b6rOJ34gwG8nlM/t2k67zMRv5f8qbayzvbcXR0DOVeF4rpw1pduLXEPmreUPs50E
Ws+xBtffhQbLf354QFdclCZlZmy1OvmO1RPJfrDNAoGBANYGf9U0Y0RspwQqKwcW
Nzb/UNAl8cQTvPPv86v67oSeZLPDh77OHuXJPWgwgaq/JP96JkF0Fd6ErJcadGGX
YLt2XzKHChVNJHOD8daFy0zd/YkfwHypb6sIqATWNXyrWYbQ2r672xGaMLEtBGgy
50l9qJuZJ2Sf2kM/I4zXUWcK
-----END PRIVATE KEY-----`

func TestGanterToken(t *testing.T) {
	priKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(privateKey))
	if err != nil {
		t.Fatalf("cannot ParseRSAPrivateKeyFromPEM priKey %v", err)
	}

	g := NewJWTTokenGen("coolcar/auth", priKey)
	g.nowFunc = func() time.Time {
		return time.Unix(1516239022, 0)
	}
	tkn, err := g.GenerateToken("630889a93396fa0232d5f3d8", 2*time.Hour)
	if err != nil {
		t.Errorf("cannot generate token: %v", err)
	}

	want := `eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MTYyNDYyMjIsImlhdCI6MTUxNjIzOTAyMiwiaXNzIjoiY29vbGNhci9hdXRoIiwic3ViIjoiNjMwODg5YTkzMzk2ZmEwMjMyZDVmM2Q4In0.hG5XVev3YeY8KtZXboHltv1XBKLMknGd82vWAVfGX9l3ejGy82Q__yAs3zPk4S2Yw3_cauT97mki_hfMDrFhdD1e2jA_wv6Eiry0RTbdLtbPe87rQ0d7wMxfkVGvlOlb7XhuvfZKD3RdU911SP1xbofWJpKcpydFUSqOl0tLGJzYoaQEICnAsK9VOEfwP7Xy00F-AphDyXssiHyHF93iXNJItuUsRaBkX3m13GVSBVlnR_Qrt2MmOsqFQ8RJJaQ1gzhnCN7vl-FMHlMLfoGSEATEmLVozls0_olcbWwxLqgXJ4AG1WFdFLLEJX2-3VvRo74u4mi-nLBDV03RqEilow`

	if tkn != want {
		t.Errorf("wrong token generated. want: %q; got: %q", want, tkn)
	}
}
