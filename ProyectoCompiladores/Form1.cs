using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Data;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Windows.Forms;
using System.IO;
using System.Text.RegularExpressions;
using System.Diagnostics;

namespace ProyectoCompiladores
{
    public partial class Form1 : Form
    {
        #region "Avoid flickering"

        private const int WM_SETREDRAW = 0xB;

        [System.Runtime.InteropServices.DllImport("User32")]

        private static extern bool SendMessage(IntPtr hWnd, int msg, int wParam, int lParam);

        private void FreezeDraw()
        { //Disable drawing
            SendMessage(codeRichTextBox.Handle, WM_SETREDRAW, 0, 0);
        }
        private void UnfreezeDraw()
        { //Enable drawing and do a redraw.
            SendMessage(codeRichTextBox.Handle, WM_SETREDRAW, 1, 0);
            codeRichTextBox.Invalidate(true);
        }
        #endregion

        OpenFileDialog ofd;
        SaveFileDialog sfd;

        String[] keywords = { "public", "void", "using", "static", "class" };
        String[] dataTypes = { "int", "bool", "char", "string" };

        HashSet<string> keywordsHashSet;
        HashSet<string> dataTypesHashSet;

        bool fileOpened;

        public Form1()
        {
            InitializeComponent();
        }

        private void Form1_Load(object sender, EventArgs e)
        {
            keywordsHashSet = new HashSet<string>(keywords);
            dataTypesHashSet = new HashSet<string>(dataTypes);
            codeRichTextBox.SelectionFont = new Font("Courier New", 10, FontStyle.Regular);
            fileNameLabel.Text = "Nuevo archivo";
            fileOpened = false;
        }

        private void openFileButton_Click(object sender, EventArgs e)
        {
            ofd = new OpenFileDialog();
            ofd.ShowDialog();
            string readFile = "";

            string fileName = ofd.FileName;
            if (fileName != "")
            {
                codeRichTextBox.Text = "";
                readFile = File.ReadAllText(ofd.FileName);
                Parse(readFile);
                fileNameLabel.Text = ofd.FileName;
                fileOpened = true;
            }          
        }

        private void saveFileButton_Click(object sender, EventArgs e)
        {
            if (fileOpened)
            {
                codeRichTextBox.SaveFile(fileNameLabel.Text, RichTextBoxStreamType.PlainText);
            }
            else
            {
                sfd = new SaveFileDialog();
                sfd.ShowDialog();
                if (sfd.FileName != "")
                {
                    codeRichTextBox.SaveFile(sfd.FileName, RichTextBoxStreamType.PlainText);
                    fileOpened = true;
                    fileNameLabel.Text = sfd.FileName;
                }
            }
        }


        void Parse(string code)
        {
            // Foreach line in input,
            // identify key words and format them when adding to the rich text box.
            Regex r = new Regex("\\n");
            string[] lines = r.Split(code);
            foreach (string l in lines)
            {
                ParseLine(l);
            }
        }

        void ParseLine(string line)
        {
            Regex r = new Regex("([ \\t{}();])");
            String[] tokens = r.Split(line);

            foreach (string token in tokens)
            {
                // Set the token's default color and font.
                codeRichTextBox.SelectionColor = Color.White;

                // Check for a comment.
                if (token == "//" || token.StartsWith("//"))
                {
                    // Find the start of the comment and then extract the whole comment.
                    int index = line.IndexOf("//");
                    string comment = line.Substring(index, line.Length - index);
                    codeRichTextBox.SelectionColor = Color.LightGreen;
                    codeRichTextBox.SelectedText = comment;
                    break;
                }

                // Check whether the token is a keyword. 

                colorToken(token);

                codeRichTextBox.SelectedText = token;
            }
            codeRichTextBox.SelectedText = "\n";
        }

        private void codeRichTextBox_TextChanged(object sender, EventArgs e)
        {
            FreezeDraw();
            // Calculate the starting position of the current line.  
            int start = 0, end = 0;
            for (start = codeRichTextBox.SelectionStart-1; start >=0; start--)
            {
                if (codeRichTextBox.Text[start] == '\n') { start++; break; }
            }
            if (start < 0) start = 0;
            // Calculate the end position of the current line.  
            for (end = codeRichTextBox.SelectionStart; end < codeRichTextBox.Text.Length; end++)
            {
                if (codeRichTextBox.Text[end] == '\n') break;
            }
            // Extract the current line that is being edited.  
            String line = codeRichTextBox.Text.Substring(start, end - start);
            
            // Backup the users current selection point.  
            int selectionStart = codeRichTextBox.SelectionStart;
            int selectionLength = codeRichTextBox.SelectionLength;

            // Split the line into tokens.  
            Regex r = new Regex("([ \\t{}();])");
            string[] tokens = r.Split(line);
            int index = start;
            foreach (string token in tokens)
            {
                codeRichTextBox.SelectionStart = index;
                codeRichTextBox.SelectionLength = token.Length;
                // Set the token's default color and font.  
                
                codeRichTextBox.SelectionColor = Color.White;

                // Check for a comment.
                if (token == "//" || token.StartsWith("//"))
                {
                    // Find the start of the comment and then extract the whole comment.
                    int length = line.Length - (index - start);
                    string commentText = codeRichTextBox.Text.Substring(index, length);
                    codeRichTextBox.SelectionStart = index;
                    codeRichTextBox.SelectionLength = length;
                    codeRichTextBox.SelectionColor = Color.LightGreen;
                    break;
                }
                // Check whether the token is a keyword. 
                colorToken(token);
                
                index+= token.Length;
                
            }
            // Restore the users current selection point.   
            codeRichTextBox.SelectionStart = selectionStart;
            codeRichTextBox.SelectionLength = selectionLength;

            UnfreezeDraw();
        }

        private void colorToken(string token){
            if (keywordsHashSet.Contains(token))
            {
                codeRichTextBox.SelectionColor = Color.Blue;
            }

            if (dataTypesHashSet.Contains(token))
            {
                codeRichTextBox.SelectionColor = Color.MediumPurple;
            }
        }

        private void compileButton_Click(object sender, EventArgs e)
        {
            Process cmd = new Process();
            cmd.StartInfo.FileName = "cmd.exe";
            cmd.StartInfo.RedirectStandardInput = true;
            cmd.StartInfo.RedirectStandardOutput = true;
            cmd.StartInfo.UseShellExecute = false;
            cmd.Start();

            cmd.StandardInput.WriteLine("cd C:/Users/Carlos/desktop");
            cmd.StandardInput.WriteLine("mkdir carpeta");
            cmd.StandardInput.WriteLine("echo 'hola'");
            cmd.StandardInput.Flush();
            cmd.StandardInput.Close();
            cmd.WaitForExit();
            Console.WriteLine(cmd.StandardOutput.ReadToEnd());

        }

        private void saveFileAsButton_Click(object sender, EventArgs e)
        {
            sfd = new SaveFileDialog();
            sfd.ShowDialog();
            if (sfd.FileName != "")
            {
                codeRichTextBox.SaveFile(sfd.FileName, RichTextBoxStreamType.PlainText);
                fileOpened = true;
                fileNameLabel.Text = sfd.FileName;
            }
        }

        private void newFileButton_Click(object sender, EventArgs e)
        {
            codeRichTextBox.Text = "";
            fileNameLabel.Text = "Nuevo archivo";
        }
    }
}
