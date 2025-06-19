
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Volume2 } from "lucide-react";

const ArabicLetters = () => {
  // Arabic alphabet data
  const arabicLetters = [
    { arabic: "ا", name: "Alif", pronunciation: "a" },
    { arabic: "ب", name: "Baa", pronunciation: "b" },
    { arabic: "ت", name: "Taa", pronunciation: "t" },
    { arabic: "ث", name: "Thaa", pronunciation: "th" },
    { arabic: "ج", name: "Jeem", pronunciation: "j" },
    { arabic: "ح", name: "Haa", pronunciation: "ḥ" },
    { arabic: "خ", name: "Khaa", pronunciation: "kh" },
    { arabic: "د", name: "Dal", pronunciation: "d" },
    { arabic: "ذ", name: "Thal", pronunciation: "dh" },
    { arabic: "ر", name: "Raa", pronunciation: "r" },
    { arabic: "ز", name: "Zay", pronunciation: "z" },
    { arabic: "س", name: "Seen", pronunciation: "s" },
    { arabic: "ش", name: "Sheen", pronunciation: "sh" },
    { arabic: "ص", name: "Sad", pronunciation: "ṣ" },
    { arabic: "ض", name: "Dad", pronunciation: "ḍ" },
    { arabic: "ط", name: "Taa", pronunciation: "ṭ" },
    { arabic: "ظ", name: "Zaa", pronunciation: "ẓ" },
    { arabic: "ع", name: "Ayn", pronunciation: "ʿ" },
    { arabic: "غ", name: "Ghayn", pronunciation: "gh" },
    { arabic: "ف", name: "Faa", pronunciation: "f" },
    { arabic: "ق", name: "Qaf", pronunciation: "q" },
    { arabic: "ك", name: "Kaf", pronunciation: "k" },
    { arabic: "ل", name: "Lam", pronunciation: "l" },
    { arabic: "م", name: "Meem", pronunciation: "m" },
    { arabic: "ن", name: "Noon", pronunciation: "n" },
    { arabic: "ه", name: "Haa", pronunciation: "h" },
    { arabic: "و", name: "Waw", pronunciation: "w" },
    { arabic: "ي", name: "Yaa", pronunciation: "y" },
  ];

  const playSound = (letter: string) => {
    // Placeholder for audio functionality
    console.log(`Playing sound for letter: ${letter}`);
  };

  return (
    <div className="space-y-8">
      <div>
        <h1 className="text-3xl font-bold text-gray-900 dark:text-gray-100 mb-2">Arabic Letters</h1>
        <p className="text-gray-600 dark:text-gray-400">Learn the 28 letters of the Arabic alphabet.</p>
      </div>

      <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 xl:grid-cols-7 gap-4">
        {arabicLetters.map((letter, index) => (
          <Card key={index} className="hover:shadow-lg transition-shadow duration-200 dark:hover:shadow-xl">
            <CardHeader className="text-center pb-3">
              <div className="text-4xl font-bold text-blue-600 dark:text-blue-400 mb-2">
                {letter.arabic}
              </div>
              <CardTitle className="text-sm">{letter.name}</CardTitle>
              <CardDescription className="text-xs">
                /{letter.pronunciation}/
              </CardDescription>
            </CardHeader>
            <CardContent className="pt-0">
              <Button
                variant="outline"
                size="sm"
                className="w-full"
                onClick={() => playSound(letter.arabic)}
              >
                <Volume2 className="h-3 w-3 mr-1" />
                Play
              </Button>
            </CardContent>
          </Card>
        ))}
      </div>

      <Card>
        <CardHeader>
          <CardTitle>About Arabic Letters</CardTitle>
        </CardHeader>
        <CardContent className="space-y-4">
          <p className="text-gray-600 dark:text-gray-400">
            The Arabic alphabet consists of 28 letters, all representing consonants. 
            Arabic is written from right to left, and letters change their shape depending 
            on their position in a word (beginning, middle, end, or isolated).
          </p>
          <div className="grid grid-cols-1 md:grid-cols-2 gap-4 mt-4">
            <div>
              <h3 className="font-semibold mb-2">Key Features:</h3>
              <ul className="text-sm text-gray-600 dark:text-gray-400 space-y-1">
                <li>• 28 letters total</li>
                <li>• Right-to-left writing</li>
                <li>• Letters connect to each other</li>
                <li>• Multiple forms per letter</li>
              </ul>
            </div>
            <div>
              <h3 className="font-semibold mb-2">Letter Types:</h3>
              <ul className="text-sm text-gray-600 dark:text-gray-400 space-y-1">
                <li>• Sun letters (14 letters)</li>
                <li>• Moon letters (14 letters)</li>
                <li>• Connecting letters</li>
                <li>• Non-connecting letters</li>
              </ul>
            </div>
          </div>
        </CardContent>
      </Card>
    </div>
  );
};

export default ArabicLetters;
