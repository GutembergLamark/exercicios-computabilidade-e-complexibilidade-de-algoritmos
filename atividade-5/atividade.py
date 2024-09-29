class AFD:
    def __init__(self):
        self.state = 'q0'
    
    def transition(self, char):
        if self.state == 'q0':  
            if char == '0':
                self.state = 'q1'
            elif char == '1':
                self.state = 'q2'
        elif self.state == 'q1':  
            if char == '1':
                self.state = 'q3'
        elif self.state == 'q2':  
            if char == '0':
                self.state = 'q4'
        elif self.state == 'q3':  
            if char == '0':
                self.state = 'q1'
        elif self.state == 'q4':  
            if char == '1':
                self.state = 'q2'
    
    def is_accepted(self, string):
        self.state = 'q0'  
        for char in string:
            self.transition(char)
        return self.state == 'q1' or self.state == 'q2'  

if __name__ == '__main__':
    afd = AFD()
    print(afd.is_accepted("0"))      
    print(afd.is_accepted("1"))      
    print(afd.is_accepted("01"))     
    print(afd.is_accepted("10"))     
    print(afd.is_accepted("1001"))   
    print(afd.is_accepted("0110"))   