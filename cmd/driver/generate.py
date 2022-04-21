import sys

def read_data(file_name):
    f = open(file_name, "r")
    data = []
    for line in f.readlines():
        line = line.strip()
        l_list = line.split()
        data.append(l_list)
    f.close()
    return data
def main():
    # usage input, output, result iat
    data = read_data(sys.argv[1])
    iat = int(sys.argv[4])
    sleep_time = [0]
    out_p = sys.argv[2]
    result_f = sys.argv[3]
    f = open(out_p,"w")
    for i in range(len(data)-1):
        sleep_time.append(iat*(int(data[i+1][3]) - int(data[i][3])))
    for i in range(len(data)):
        f.write("sleep {}\n".format(sleep_time[i]/1000))
        # one_exec.go pred id para output_file
        if sys.argv[-1] == "0":
            f.write("./main -p {} -i {} -a {} -o {} &\n".format(0, data[i][0], data[i][2], result_f))
        else:
            f.write("./main -p {} -i {} -a {} -o {} &\n".format(data[i][2], data[i][0], data[i][2], result_f))
    f.close()

if __name__ == "__main__":
    main()

