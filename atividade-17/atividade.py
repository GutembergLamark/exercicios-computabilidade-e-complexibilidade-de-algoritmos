class AFN:
    def __init__(self):
        self.states = {'q0', 'q1', 'q2'}
        self.start_state = 'q0'
        self.accept_state = 'q2'
        self.transitions = {
            'q0': {'0': {'q1'}, '1': {'q0'}},
            'q1': {'0': {'q1'}, '1': {'q2'}},
            'q2': {'0': {'q1'}, '1': {'q0'}}
        }

    def is_accepted(self, string):
        current_states = {self.start_state}  

        for char in string:
            next_states = set()
            for state in current_states:
                if char in self.transitions[state]:
                    next_states.update(self.transitions[state][char])
            current_states = next_states

        return self.accept_state in current_states
    
class AFD:
    def __init__(self):
        self.start_state = frozenset({'q0'})
        self.accept_state = frozenset({'q2'})
        self.transitions = {
            frozenset({'q0'}): {'0': frozenset({'q1'}), '1': frozenset({'q0'})},
            frozenset({'q1'}): {'0': frozenset({'q1'}), '1': frozenset({'q2'})},
            frozenset({'q2'}): {'0': frozenset({'q1'}), '1': frozenset({'q0'})},
        }

    def is_accepted(self, string):
        current_state = self.start_state  

        for char in string:
            if char in self.transitions[current_state]:
                current_state = self.transitions[current_state][char]
            else:
                return False

        return current_state == self.accept_state

if __name__ == '__main__':
  print('AFN \n')
  afn = AFN()
  print(afn.is_accepted("1101"))  
  print(afn.is_accepted("00101"))  
  print(afn.is_accepted("010"))  


  print('\nAFD \n')
  afd = AFD()
  print(afd.is_accepted("1101"))   
  print(afd.is_accepted("00101"))  
  print(afd.is_accepted("010"))    
  print(afd.is_accepted("101"))    
  print(afd.is_accepted("011"))    