def read_file(file_path):
    fl = open(file_path)
    a = fl.readlines()
    return a


def main():
    cust = []
    prd = []
    combined = []
    customers = read_file("csv/lista1.csv")
    products = read_file("csv/lista2.csv")
    for customer in customers[1:]:
        splt = customer.split(";")
        id = splt[0]
        cname = splt[1].replace("\n", "")
        cust.append((id, cname))
    for product in products[1:]:
        splt = product.split(";")
        cu = splt[0]
        prid = splt[1]
        price = splt[2]
        cur = splt[3].replace("\n", "")
        prd.append((cu, prid, price, cur))
    for cus in cust:
        for pr in prd:
            if cus[1] == pr[0]:
                combined.append(("", pr[3], cus[0], pr[1], "",  pr[2]))
    ou = "Version*;Currency ID*;Customer ID*;Product ID*;Key Figure*;Price\n"
    for com in combined:
        ou += ";".join(com) + "\n"
    out_file = open("output.csv", "w+")
    out_file.write(ou)

if __name__ == "__main__":
    main()
