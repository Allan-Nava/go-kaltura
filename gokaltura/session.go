package gokaltura

/*
curl -X POST https://www.kaltura.com/api_v3/service/session/action/start \
    -d "secret=********" \
    -d "userId=YOUR_USER_ID" \
    -d "type=0" \
    -d "expiry=86400"
*/



/*
curl -X POST https://www.kaltura.com/api_v3/service/session/action/end \
    -d "ks=$KALTURA_SESSION" \
*/


/*
curl -X POST https://www.kaltura.com/api_v3/service/session/action/get \
    -d "ks=$KALTURA_SESSION" \
*/



/*
curl -X POST https://www.kaltura.com/api_v3/service/session/action/impersonate \
    -d "secret=********" \
    -d "userId=YOUR_USER_ID" \
    -d "type=0" \
    -d "expiry=86400"
*/

/*
curl -X POST https://www.kaltura.com/api_v3/service/session/action/startWidgetSession \
    -d "expiry=86400"
*/

/*
curl -X POST https://www.kaltura.com/api_v3/service/session/action/impersonateByKs \
    -d "ks=$KALTURA_SESSION" \
*/