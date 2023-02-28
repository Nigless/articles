import Block from "./Block";
import Link from "./Link";
import ThemeSwitcher from "./ThemeSwitcher";

function App() {
  return (
    <>
      <div class="m-auto max-w-3xl px-4 sm:px-8">
        <div class="flex text-action-main h-16 items-center justify-between">
          Article v0.0.0
          <ThemeSwitcher />
        </div>
        <div class="mt-32">
          <h1 class="text-4xl font-bold">What is CRISPR-Cas9?</h1>
          <div class="text-action-main text-sm space-x-3 mt-3">
            <Link href="#" color="neutral">
              jasper Maglot
            </Link>
            <span>
              Feb 14
            </span>
            <span>
              15 min read
            </span>
          </div>
        </div>
        <div class="markdown mb-64">
          <blockquote>
            CRISPR-Cas9 is a genome editing tool that is creating a buzz in the science world. It is faster, cheaper and more accurate than previous techniques of editing DNA and has a wide range of potential applications.
          </blockquote>
          <ul>
            <li>
              <code>CRISPR-Cas9</code> is a unique technology that enables geneticists and medical researchers to edit parts of the genome by removing, adding or altering sections of the DNA sequence.
            </li>
            <li>
              It is currently the simplest, most versatile and precise method of genetic manipulation and is therefore causing a buzz in the science world.
            </li>
          </ul>
          <h2>
            How does it work?
          </h2>
          <ul>
            <li>
              The <a href="#">CRISPR-Cas9</a> system consists of two key molecules that introduce a change (mutation) into the DNA. These are:

              <ul>
                <li>
                  an enzyme called Cas9. This acts as a pair of ‘molecular scissors’ that can cut the two strands of DNA at a specific location in the genome so that bits of DNA can then be added or removed.
                </li>
                <li>
                  a piece of RNA called guide RNA (gRNA). This consists of a small piece of pre-designed RNA sequence (about 20 bases long) located within a longer RNA scaffold. The scaffold part binds to DNA and the pre-designed sequence ‘guides’ Cas9 to the right part of the genome. This makes sure that the Cas9 enzyme cuts at the right point in the genome.
                </li>
              </ul>
            </li>
            <li>
              The guide RNA is designed to find and bind to a specific sequence in the DNA. The guide RNA has RNA bases that are complementary to those of the target DNA sequence in the genome. This means that, at least in theory, the guide RNA will only bind to the target sequence and no other regions of the genome.
            </li>
            <li>
              The Cas9 follows the guide RNA to the same location in the DNA sequence and makes a cut across both strands of the DNA.
            </li>
            <li>
              At this stage the cell recognises that the DNA is damaged and tries to repair it.
            </li>
            <li>
              Scientists can use the DNA repair machinery to introduce changes to one or more genes in the genome of a cell of interest.
            </li>
          </ul>
        </div>
      </div>
    </>
  );
};

export default App;
