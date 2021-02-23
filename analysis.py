import os, time

# which data set to be used
DataSet = "p2"
# output filename
now = time.strftime("%Y-%m-%d %H:%M:%S", time.localtime())
OutputFileName = "result" + "-" + DataSet + "-" + now + ".csv"
# number of repeat in order to calculate average run time
repeat = 10
# upper limit of threads
maxThread = 128
# lower limit of threads
minThread = 1

def execCmd(cmd):
    r = os.popen(cmd)
    text = r.read()
    r.close()
    return text

def writeFile(filename, data):
    f = open(filename, "w")
    f.write(data)
    f.close()

if __name__ == "__main__":
    csv = ""
    for threads in range(minThread, maxThread+1):
        unit = "ms"
        totalRuntime = 0
        exclude = 0
        for r in range(repeat): 
            cmd = "go run ./ " + DataSet + " " + str(threads)
            result = execCmd(cmd)
            result = result.strip()
            splt = result.split("runtime: ")
            runTime = splt[1]
            time = float(runTime[:-2])
            unit2 = runTime[-2:]
            print("thread: " + str(threads) + " repeat: " + str(r) + runTime )
            if unit2==unit :
                totalRuntime += time
            else :
                print("unit is different! exclude: " + str(time) + str(unit2))
                exclude += 1
        if repeat == exclude:
            print("too many excludes, analysis fail")
            exit()
        averageRuntime = totalRuntime / (repeat - exclude)
        csv += str(threads) + "," + str(averageRuntime) + "\n"
    writeFile(OutputFileName, csv)
            


