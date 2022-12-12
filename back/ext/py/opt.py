def main():
    i = 0
    arrMsg = []
    arrAns = []

    while i < 4:
        m = str(input("Message > "))
        arrMsg.append(m)

        a = str(input("Answer  > "))
        arrAns.append(a)

        i += 1
    
    for k, message in enumerate(arrMsg):
        m = message
        a = arrAns[k]
        t = '"message" : "{}", "answer" : "{}"'
        t = t.format(m, a)
        
        t_final = "{" + t + "},"
        print(t_final)
main()
