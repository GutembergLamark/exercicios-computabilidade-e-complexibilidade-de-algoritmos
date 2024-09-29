class AFD:
    def __init__(self):
        self.state = 'q0'  

    def transition(self, char):
        if self.state == 'q0':
            if char == '1':
                self.state = 'q1'
            elif char == '0':
                self.state = 'q-1' 
        elif self.state == 'q1':
            if char == '1':
                self.state = 'q2' 
            elif char == '0':
                self.state = 'q0'
        elif self.state == 'q-1':
            if char == '1':
                self.state = 'q0' 
            elif char == '0':
                self.state = 'q-2' 
        elif self.state == 'q2':
            if char == '1':
                self.state = 'q3'
            elif char == '0':
                self.state = 'q1' 
        elif self.state == 'q-2':
            if char == '1':
                self.state = 'q-1' 
            elif char == '0':
                self.state = 'q-3' 

    def is_accepted(self, string):
        self.state = 'q0' 
        for char in string:
            self.transition(char)

        return self.state in {'q1', 'q2', 'q3'}

if __name__ == '__main__':
    afd = AFD()
    print(afd.is_accepted("110"))   
    print(afd.is_accepted("101"))   
    print(afd.is_accepted("11100"))
    print(afd.is_accepted("100"))   
    print(afd.is_accepted("111"))   