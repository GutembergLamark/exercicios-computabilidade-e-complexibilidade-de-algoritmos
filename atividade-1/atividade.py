class AFD:
    def __init__(self):
        self.state = 'q0'
    
    def transition(self, char):
        if self.state == 'q0':
            if char == '0':
                self.state = 'q0'
            elif char == '1':
                self.state = 'q1'
        elif self.state == 'q1':
            if char == '0':
                self.state = 'q0'
            elif char == '1':
                self.state = 'q1'
    
    def is_accepted(self, string):
        self.state = 'q0' 
        for char in string:
            self.transition(char)
        return self.state == 'q1'

if __name__ == '__main__':
    afd = AFD()
    print(afd.is_accepted("110"))  
    print(afd.is_accepted("111"))  
    print(afd.is_accepted("101"))  
    print(afd.is_accepted("0"))    