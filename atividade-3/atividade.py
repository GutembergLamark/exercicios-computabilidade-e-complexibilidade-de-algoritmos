class AFD:
    def __init__(self):
        self.state = 'q0'
    
    def transition(self, char):
        if self.state == 'q0':
            if char == '1':
                self.state = 'q1'
        elif self.state == 'q1':
            if char == '1':
                self.state = 'q2'
        elif self.state == 'q2':
            if char == '1':
                self.state = 'q3'
    
    def is_accepted(self, string):
        self.state = 'q0'
        for char in string:
            self.transition(char)
        return self.state == 'q2'

if __name__ == '__main__':
    afd = AFD()
    print(afd.is_accepted("110"))  
    print(afd.is_accepted("101"))  
    print(afd.is_accepted("1001"))
    print(afd.is_accepted("01"))    
    print(afd.is_accepted("00111"))

# Entendi da questão que a string não precisa conter apenas dois '1's, mas sim que a quantidade de '1's precisa ser igual a 2 