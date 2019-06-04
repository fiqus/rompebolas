# Rompebolas

a simple tool to remind developers to log their work at the end of the day

when run, it opens in full screen and prompt for a log message

the message is saved, with the date it was generated on, in a filed called ~/rompe.txt

## Recommended install instructions

1.  get dependencies
        
        make setup
        
2. build the project

        make build
        
3. copy the executable somewhere you like. eg:

        cp rompebolas ~/bin/
        
4. create a crontab entry to run it whenever you want. eg:

        crontab -e
        
        # insert the following line in the crontab, with the correct path to the executable
        
        [[minutes]] [[hour]] * * 1-5 DISPLAY=:0.0 [[EXECUTABLE_PATH]]rompebolas >/dev/null 2>&1
        
        # eg:
        
        50 17 * * 1-5 DISPLAY=:0.0 /home/mlambir/bin/rompebolas >/dev/null 2>&1

this will run the program once a day every weekday at 5:50 pm