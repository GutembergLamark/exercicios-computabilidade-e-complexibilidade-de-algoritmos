class AFN:
    def __init__(self):
        self.states = {'q0', 'q1', 'q2'}  
        self.start_state = 'q0'
        self.accept_states = {'q0', 'q1', 'q2'}

    def is_accepted(self, string):
        current_states = {self.start_state}

        for char in string:
            next_states = set()
            for state in current_states:
                if state == 'q0':
                    if char == '0':
                        next_states.add('q1')  
                    elif char == '1':
                        next_states.add('q2')  
                elif state == 'q1':  
                    if char == '1':
                        next_states.add('q2')  
                elif state == 'q2':  
                    if char == '0':
                        next_states.add('q1')  
            current_states = next_states

        
        return len(current_states) > 0

if __name__ == '__main__':
    afn = AFN()
    print(afn.is_accepted("010101"))
    print(afn.is_accepted("101010"))
    print(afn.is_accepted("1101"))  
    print(afn.is_accepted("100"))   
    print(afn.is_accepted(""))