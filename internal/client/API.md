# Draw

```
POST https://backend.wplace.live/s0/pixel/<canvas_x>/<canvas_y>
COOKIES cf_clearance, j
REQ {colors:[c1,c2,...,cn],coords:[x1,y1,x2,y2,...,xn,yn],t:"???"}
RES {painted:n}
```

# Query Canvas

```
GET https://backend.wplace.live/files/s0/tiles/<canvas_x>/<canvas_y>.png
RES image/png 1000x1000
```

# User Data

```
GET https://backend.wplace.live/me
COOKIES cf_clearance, j
RES:
{
  "allianceId": uint32,
  "allianceRole": string,
  "charges": {
    "cooldownMs": uint32,
    "count": double,
    "max": uint32
  },
  "country": string,
  "discord": string,
  "droplets": uint32,
  "equippedFlag": uint32,
  "extraColorsBitmap": uint32,
  "favoriteLocations": [
    {
      "id": uint32,
      "name": string,
      "latitude": double,
      "longitude": double
    },
    ...
  ],
  "flagsBitmap": string,
  "id": uint32,
  "isCustomer": bool,
  "level": double,
  "maxFavoriteLocations": uint32,
  "name": string,
  "needsPhoneVerification": bool,
  "picture": string,
  "pixelsPainted": uint32,
  "showLastPixel": bool
}
```