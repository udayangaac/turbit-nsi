# Turbit Notification Selector & Indexer

[![CircleCI](https://circleci.com/gh/udayangaac/turbit-nsi.svg?style=svg)](https://circleci.com/gh/udayangaac/turbit-nsi)
<br/>

This Service responsible for getting notification list based on location (Phase one)
######  Overview
![Overall View](/docs/overall_view_phase_one.png)

##### APIs

######  Add Notification

- Request
```
PUT /tnsi/notification HTTP/1.1
Host: localhost:3001
Content-Type: application/json

{
	"id": <number>,
    "company_name": "Turbit ",
    "content": "This is a test document",
    "notification_type": 1,
    "start_time": <time stamp>,
    "end_date": <time stamp>,
    "logo_company": <image_url>,
    "image_publisher": <image_url>,
    "category": "<category_name>",
    "locations": [
        {
            "lat": "6.714360",
            "lon": "81.059219"
        },
        {
            "lat": "6.814360",
            "lon": "81.059219"
        }
    ]
}

```

- Response
```
{
    "data": {
        "message": "Added notification successfully !"
    }
}
```


######  Modify Notification

- Request
```
PUT /tnsi/notification/{notification_id} HTTP/1.1
Host: localhost:3001
Content-Type: application/json

{
    "company_name": "Turbit ",
    "content": "This is a test document",
    "notification_type": 1,
    "start_time": <time stamp>,
    "end_date": <time stamp>,
    "logo_company": <image_url>,
    "image_publisher": <image_url>,
    "categories": ["kfc","food"],
    "locations": [
        {
            "lat": "6.714360",
            "lon": "81.059219"
        },
        {
            "lat": "6.814360",
            "lon": "81.059219"
        }
    ]
}

```

- Response
```
{
    "data": {
        "message": "Modified notification successfully !"
    }
}
```

######  Get Notifications

- Request
```
POST /tnsi/notifications HTTP/1.1
Host: localhost:3001
Content-Type: application/json

{
	"lat": "6.814360",
    "lon": "81.059219",
    "geo_ref_id":<geo_ref_id>, //optional
    "user_id":<id>
    "categories":["food"],
    "is_newest": false
}
```

- Response
```
{
    {
        "data": {
            "geo_ref_id": "0x88610250c5fffff_3",
            "notification": [
                {
                    "id": 1,
                    "company_name": "KFC",
                    "content": "Offer !",
                    "notification_type": 1,
                    "start_time": "2020-04-21T19:38:07+00:00",
                    "end_date": "2020-10-21T19:38:07+00:00",
                    "logo_company": "https://upload.wikimedia.org/wikipedia/en/thumb/b/bf/KFC_logo.svg/1024px-KFC_logo.svg.png",
                    "image_publisher": "https://i.pinimg.com/736x/98/90/85/9890854f31eebc125f5e1c2956e132c7.jpg",
                    "categories": [
                        "food"
                    ],
                    "locations": [
                        {
                            "lat": "6.814360",
                            "lon": "81.059219"
                        }
                    ],
                    "geo_hex_ids": [
                        "0x88610250c5fffff"
                    ]
                }
            ]
        }
    }
}
```


