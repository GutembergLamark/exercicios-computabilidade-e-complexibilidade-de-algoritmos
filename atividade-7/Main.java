import java.util.regex.Matcher;
import java.util.regex.Pattern;

class AFN {
    public static boolean is_accepted(String string) {
        Pattern pattern = Pattern.compile("^01.*10$");
        Matcher matcher = pattern.matcher(string);
        return matcher.find();
    }
}

public class Main {

    public static void main(String[] args) {
        AFN afn = new AFN();
        System.out.println(afn.is_accepted("0110"));  
        System.out.println(afn.is_accepted("010110")); 
        System.out.println(afn.is_accepted("10110")); 
        System.out.println(afn.is_accepted("0111"));  
    }
    
}