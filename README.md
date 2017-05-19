# logger
A simple logger with levels

Example:

   import "logger"

   log := GetLogger(INFO, "logger")

   log.Info("should print")

   log.Debug("current level is %v", log.Level) // but this will not print

   log.Fatal("print and die!")
