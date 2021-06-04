package packed

import "github.com/gogf/gf/os/gres"

func init() {
	if err := gres.Add("H4sIAAAAAAAC/wrwZmYRYeBg4GDQdDkSxIAERBg4GYrLE9PTU4v0obReVnF+XmgIKwOj/9OdCVvCYv2v+Ijsex5uUnuf22Pa3abyvY/dFoezHZ46KeDW5tQJ56bMXNkywbdBg6kpwUF4YgOvwM4kVanAPw+O75gb+7/15fFDZb8Pzroal3PJ+LnN5zOfz1je3SnnOUll5kvhBMbJoS/rBDiahOfcLvpgFdVxKc3kMovW6kk2ESFPNLoKZp5YZVIemslwWXOTpdKN87K7XfOCzjYZ5v/f8+5c9f1ne3vFD8UKVc9/1f58q/z5I+92bX/Fd3ev7jz5Smn2R1eDH8r9mRPfU8ZhYPkoWGDBzKd1zDpG18Juz1ojfMZ07u81WroP3R0ij2QuWNGS3n0tNfqxsE5dz6TDDdMkjAr3Jx+fdojrUIWS1LqAJIuq2fM28pzdtnzCup9cU9UkDxYp5W6ZcOTprJnNkfftxE6nL5Sb+9Dw+CYTq9evBUXrLb5Nu+Yk/mQKU/Tn2icOmzNVnmbYvb2udGdNoK3w8m7N2bHTEs9y9bCfEZx/9rQL4zsLw6LNGj7WHU2Kq4WW8z1JYctduLRK5vTv+U437//cP4kxtWVivVPYpg7G4KOCbCUn3L5P6HR/kzjDz+pyjdZK7ZMZsyaLzPA1YJ7mGSYXMuXOZJU3kZXXC66/it9+5F7evDl3O6+oxO05zCHXF7fruobczqrVR6r2tXrs+Wy8xruHwS6bPXxyj3/Tudshn4438E5ZzBOe0vOl/eLTcx/WdJ63cL3hcyrI823gsRyny+6Tr/C+4ll4bL+j7Bsms9Mvp+2cX3fy36bZVoZFc3Ky18XfXi99NevvW6f89Yu6nvrml+5X4F1x72Hm48eBjlUPNE7xSgToX/71yyjqwySTQzJ/8ovdd0lLnX93o2xdb3fu+7PzKn7GrZn/5Zrd+4rVeT9aP/iVZcWH7VFP3/o++mGRga5c/GlToR3Kq8QMls2cq61tZvbXnk18XcSKuOOP+9+H+P+I/inLwPD/f4A3OwdTlGD3aiYGBnNOBgZYqmdg0EZL9eyIVA9O6MFPdyaAdCOrCfBmZBJhRuQaZJNBuQYGtjWCSLx5CGEUdqdAgADDf8dvTAxYHMbKBpJnYmBi6GRgYDBhBvEAAQAA//+np5aX0wMAAA=="); err != nil {
		panic("add binary content to resource manager failed: " + err.Error())
	}
}
