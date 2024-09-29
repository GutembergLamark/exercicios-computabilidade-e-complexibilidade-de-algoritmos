import re

class AFN:
    def is_accepted(self, string):
        padrao = r"101|110"
        return bool(re.search(padrao, string))

if __name__ == '__main__':
    afn = AFN()
    print(afn.is_accepted("101"))    # True
    print(afn.is_accepted("110"))    # True
    print(afn.is_accepted("0101"))   # True
    print(afn.is_accepted("0110"))   # True
    print(afn.is_accepted("1110"))   # True
    print(afn.is_accepted("000"))    # False
    print(afn.is_accepted("001"))    # False