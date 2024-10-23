#include "kernel/types.h"
#include "kernel/param.h"
#include "user/user.h"

int main(int argc, char *argv[])
{
	if (argc < 2)
	{
		fprintf(2, "Usage: xargs command [args...]\n");
		exit(1);
	}
	// args = {echo, bye, hello, too, 0}
	char buffer[512];
	char *args[MAXARG];
	int i;
	for (i = 1; i < argc && i < MAXARG - 1; i++)
	{
		args[i - 1] = argv[i];
	}
	// args = {echo, bye, buffer=hello too, 0}
	int len = 0;
	char c;
	int r;
	while ((r = read(0, &c, 1)) > 0)
	{
		len = 0;
		do
		{
			if (c == '\n')
				break;
			buffer[len++] = c;

			if (len >= sizeof(buffer) - 1)
			{
				fprintf(2, "xargs: input line too long\n");
				exit(1);
			}
		} while ((r = read(0, &c, 1)) > 0);
		buffer[len] = '\0';
		//buffer = "hello too"
		args[i - 1] = buffer;
		args[i] = 0;
		if (fork() == 0)
		{
			exec(args[0], args); //auto exit
			fprintf(2, "xargs: execute %s failed\n", args[0]);
			exit(1);
		}
		else
		{
			wait(0);
		}
	}
	if (r < 0)
	{
		fprintf(2, "xargs: error reading input\n");
	}
	exit(0);
}
