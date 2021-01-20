package conf;

import java.io.FileReader;
import java.io.FileWriter;
import java.io.IOException;
import java.util.HashMap;
import java.util.Map;
import java.util.Properties;

public class Config {

    public static final String CONFIG_FILE = "list-files.properties";
    private static Properties properties = new Properties();

    public static String DATASET = "/data";
    public static String EXTENSIONS = "java";
    public static Map<String, Boolean> EXTENSION_MAP = new HashMap<>();
    public static String OUTPUT = "files.txt";

    public static void save() throws IOException {
        properties.setProperty("dataset", DATASET);
        properties.setProperty("extensions", EXTENSIONS);
        properties.setProperty("output", OUTPUT);
        FileWriter writer = new FileWriter(CONFIG_FILE);
        properties.store(writer, "");
        writer.close();
    }

    public static void load()throws IOException{
        FileReader reader = new FileReader(CONFIG_FILE);
        properties.load(reader);
        reader.close();
        DATASET = properties.getProperty("dataset");
        EXTENSIONS = properties.getProperty("extensions");
        for (String ext: EXTENSIONS.split(",")){
            EXTENSION_MAP.put(ext, true);
        }
        OUTPUT = properties.getProperty("output");
    }
}
