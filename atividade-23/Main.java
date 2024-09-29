import java.util.regex.Matcher;
import java.util.regex.Pattern;

class AFN {
    public static boolean isAccepted(String string) {
        String pattern = "101|110"; 
        Pattern compiledPattern = Pattern.compile(pattern);
        Matcher matcher = compiledPattern.matcher(string);
        return matcher.find();  
    }

    
}

public class Main {

   public static void main(String[] args) {
        AFN afn = new AFN();
        System.out.println(afn.isAccepted("101"));   
        System.out.println(afn.isAccepted("110"));   
        System.out.println(afn.isAccepted("0101"));  
        System.out.println(afn.isAccepted("0110"));  
        System.out.println(afn.isAccepted("1110"));  
        System.out.println(afn.isAccepted("000"));   
        System.out.println(afn.isAccepted("001"));   
    }
}