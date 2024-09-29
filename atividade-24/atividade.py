class AFD:
    def __init__(self):
        self.state = 'q0'
        self.accept_states = {'q3'}
        self.rotate = 0

    def transition(self, char):
        if self.state == 'q0':
            if char == '0':
                self.state = 'q1'
        elif self.state == 'q1':
            if char == '1':
                self.state = 'q2'
            else:
                self.state = 'q0'
        elif self.state == 'q2':
            if char == '0' and self.rotate == 0:
                self.state = 'q0'
                self.rotate = 1
            elif char == '0' and self.rotate == 1:
                self.state = 'q3'
            elif char == '1':
                self.state = 'q0'
        
        

    def is_accepted(self, string):
        self.state = 'q0'
        self.rotate = 0
        for char in string:
            self.transition(char)
        return self.state in self.accept_states

if __name__ == '__main__':
    afd = AFD()
    print(afd.is_accepted("010010"))  
    print(afd.is_accepted("010"))      
    print(afd.is_accepted("0101010"))  
    print(afd.is_accepted("101010"))   
    print(afd.is_accepted("111"))      
    print(afd.is_accepted("000"))      