class AFD {
    private String state;

    public AFD() {
        this.state = "q0";
    }

    public void transition(char c) {
        switch (state) {
            case "q0":
                if (c == '1') state = "q1";
                break;
            case "q1":
                if (c == '1') state = "q2";
                break;
            case "q2":
                if (c == '1') state = "q3";
                break;
        }
    }

    public boolean isAccepted(String input) {
        state = "q0"; 
        for (char c : input.toCharArray()) {
            transition(c);
        }
        return "q2".equals(state); 
    }
}

public class Main {

    public static void main(String[] args) {
        AFD afd = new AFD();
        System.out.println(afd.isAccepted("110"));    
        System.out.println(afd.isAccepted("101"));    
        System.out.println(afd.isAccepted("1001"));   
        System.out.println(afd.isAccepted("01"));      
        System.out.println(afd.isAccepted("00"));     
    }
    
}

// Entendi da questão que a string não precisa conter apenas dois '1's, mas sim que a quantidade de '1's precisa ser igual a 2 