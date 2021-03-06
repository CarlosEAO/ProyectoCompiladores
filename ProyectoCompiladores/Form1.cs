﻿using System;
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
using Newtonsoft.Json;

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

        String[] keywords = { "program", "if","then", "else", "fi", "do", "until", "while", "read", "write", "not","and","or" };
        String[] dataTypes = { "float", "bool", "int" };

        HashSet<string> keywordsHashSet;
        HashSet<string> dataTypesHashSet;

        bool fileOpened;
        bool fileSaved;

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
            fileSaved = false;
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
                fileSaved = true;
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
            Regex r = new Regex("([ \\t{}();,\\n])");
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

            fileSaved = false;

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
            Regex r = new Regex("([ \\t{}();,\\n])");
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
                codeRichTextBox.SelectionColor = Color.Red;
            }

            if (dataTypesHashSet.Contains(token))
            {
                codeRichTextBox.SelectionColor = Color.Yellow;
            }
            if ( (token.Substring(0) == "<" || token.Substring(0).Equals('"')) && (token.Substring(token.Length) == ">" || token.Substring(token.Length).Equals('"')) )
            {
                codeRichTextBox.SelectionColor = Color.Orange;
            }
        }

        private void saveFileAs()
        {
            sfd = new SaveFileDialog();
            sfd.ShowDialog();
            if (sfd.FileName != "")
            {
                codeRichTextBox.SaveFile(sfd.FileName, RichTextBoxStreamType.PlainText);
                fileOpened = true;
                fileNameLabel.Text = sfd.FileName;            
                fileSaved = true;
            }
        }
        private void saveFile()
        {
            if (fileOpened)
            {
                codeRichTextBox.SaveFile(fileNameLabel.Text, RichTextBoxStreamType.PlainText);
                fileSaved = true;
            }
            else
            {
                saveFileAs();
            }
        }

        private void saveFileButton_Click(object sender, EventArgs e)
        {
            saveFile();   
        }

        private void saveFileAsButton_Click(object sender, EventArgs e)
        {
            saveFileAs();
        }

        private void newFileButton_Click(object sender, EventArgs e)
        {
            DialogResult dialogResult = MessageBox.Show("¿Quieres guardar el archivo actual?", "...", MessageBoxButtons.YesNoCancel);
            
           
            if (dialogResult == DialogResult.Yes)
            {
                saveFile();
            }
            if (dialogResult == DialogResult.No || dialogResult == DialogResult.Yes)
            {
                codeRichTextBox.Text = "";
                fileNameLabel.Text = "Nuevo archivo";
                fileOpened = false;
                fileSaved = false;
            }
           
        }

        private void recurreSyntaxTree(NodeSyntaxTreeView currentViewNode, NodeSyntaxTree currentNode)
        {
            
            currentViewNode.ProductionName = currentNode.ProductionName;
            
            currentViewNode.Token = currentNode.Token;

            currentViewNode.Text = currentViewNode.ToString();

            if (currentNode.Childs == null)
            {
                return;
            }

            foreach (NodeSyntaxTree child in currentNode.Childs)
            {
                NodeSyntaxTreeView newNode = new NodeSyntaxTreeView();
                currentViewNode.Nodes.Add(newNode);
                recurreSyntaxTree(newNode, child);
                
            }
        }


        private void recurreAttributedTree(AttributedNodeSyntaxTreeView currentViewNode, AttributedNodeSyntaxTree currentNode)
        {

            currentViewNode.ProductionName = currentNode.ProductionName;

            currentViewNode.Token = currentNode.Token;
            currentViewNode.AttributesAsString = string.Join(Environment.NewLine, currentNode.Attributes);

            currentViewNode.Text = currentViewNode.ToString();

            if (currentNode.Childs == null)
            {
                return;
            }

            foreach (AttributedNodeSyntaxTree child in currentNode.Childs)
            {
                AttributedNodeSyntaxTreeView newNode = new AttributedNodeSyntaxTreeView();
                currentViewNode.Nodes.Add(newNode);
                recurreAttributedTree(newNode, child);

            }
        }

        private void compileButton_Click(object sender, EventArgs e)
        {
            saveFile();
            if (!fileSaved) return;

            //Lexic

            Process lexicProcess = new Process();
            lexicProcess.StartInfo.WorkingDirectory = Path.Combine(Directory.GetParent(Environment.CurrentDirectory).Parent.FullName, "compiler");
            lexicProcess.StartInfo.FileName = "lexic.exe";
            lexicProcess.StartInfo.Arguments = fileNameLabel.Text;
            lexicProcess.Start();
            lexicProcess.WaitForExit();

            string lexicOutput = File.ReadAllText(Path.Combine(Directory.GetParent(Environment.CurrentDirectory).Parent.FullName, "compiler\\output\\tokens.txt"));

            //Syntax

            Process syntaxProcess = new Process();
            syntaxProcess.StartInfo.WorkingDirectory = Path.Combine(Directory.GetParent(Environment.CurrentDirectory).Parent.FullName, "compiler");
            syntaxProcess.StartInfo.FileName = "syntaxTree.exe";
            syntaxProcess.StartInfo.Arguments = Path.Combine(Directory.GetParent(Environment.CurrentDirectory).Parent.FullName, "compiler\\output\\tokens.txt");
            syntaxProcess.Start();
            syntaxProcess.WaitForExit();

            string syntaxOutput = File.ReadAllText(Path.Combine(Directory.GetParent(Environment.CurrentDirectory).Parent.FullName, "compiler\\output\\parseTree.txt"));

            NodeSyntaxTree parseTreeRoot = JsonConvert.DeserializeObject<NodeSyntaxTree>(syntaxOutput);

            NodeSyntaxTreeView parseTreeViewRoot = new NodeSyntaxTreeView();

            recurreSyntaxTree(parseTreeViewRoot, parseTreeRoot);
            parseTreeView.Nodes.Clear();
            parseTreeView.Nodes.Add(parseTreeViewRoot);

            //Semantic
            
            Process semanticProcess = new Process();
            semanticProcess.StartInfo.WorkingDirectory = Path.Combine(Directory.GetParent(Environment.CurrentDirectory).Parent.FullName, "compiler");
            semanticProcess.StartInfo.FileName = "semantic.exe";
            semanticProcess.StartInfo.Arguments = Path.Combine(Directory.GetParent(Environment.CurrentDirectory).Parent.FullName, "compiler\\output\\parseTree.txt");
            semanticProcess.Start();
            semanticProcess.WaitForExit();

            string semanticOutput = File.ReadAllText(Path.Combine(Directory.GetParent(Environment.CurrentDirectory).Parent.FullName, "compiler\\output\\attributedST.txt"));

            AttributedNodeSyntaxTree attributedTreeRoot = JsonConvert.DeserializeObject<AttributedNodeSyntaxTree>(semanticOutput);

            AttributedNodeSyntaxTreeView attributedTreeViewRoot = new AttributedNodeSyntaxTreeView();

            recurreAttributedTree(attributedTreeViewRoot, attributedTreeRoot);
            attributedTreeView.Nodes.Clear();
            attributedTreeView.Nodes.Add(attributedTreeViewRoot);

            //intermediateCode

            Process intermediateCodeProcess = new Process();
            intermediateCodeProcess.StartInfo.WorkingDirectory = Path.Combine(Directory.GetParent(Environment.CurrentDirectory).Parent.FullName, "compiler");
            intermediateCodeProcess.StartInfo.FileName = "intermediateCode.exe";
            intermediateCodeProcess.StartInfo.Arguments = Path.Combine(Directory.GetParent(Environment.CurrentDirectory).Parent.FullName, "compiler\\output\\attributedST.txt");
            intermediateCodeProcess.Start();
            intermediateCodeProcess.WaitForExit();

            string intermediateCodeOutput = File.ReadAllText(Path.Combine(Directory.GetParent(Environment.CurrentDirectory).Parent.FullName, "compiler\\output\\intermediateCode.txt"));

            intermediateCodeRichTextBox.Text = intermediateCodeOutput;
            errorsRichTextBox.Text = File.ReadAllText(Path.Combine(Directory.GetParent(Environment.CurrentDirectory).Parent.FullName, "compiler\\output\\errors.txt"));

        }
    }

    public class Token
    {
        public int Type { get; set; }
        public string Name { get; set; }
        public string Lexeme { get; set; }
        public int Row { get; set; }
        public int Column { get; set; }

        public override string ToString()
        {
            return "Name:" +Name.ToString() + " Lexeme: " + Lexeme.ToString() + ". \n";
        }

    }

    //USADA PARA DESERIALIZAR EL JSON QUE ARROJA EL PROCESO

    public class NodeSyntaxTree
    {
        public string ProductionName { get; set; }
        public Token Token { get; set; }
        public NodeSyntaxTree[] Childs { get; set; }
    }

    //USADA PARA POPULAR LOS CAMPOS DEL TREEVIEW CORRESPONDIENTE

    public class NodeSyntaxTreeView : TreeNode
    {
        public string ProductionName { get; set; }
        public Token Token { get; set; }
        public override string ToString()
        {
            if (Token.Type == 0)
            {
                return "Production Name: " + ProductionName.ToString() + ". \n";
            }

            return "Token: \n" + Token.ToString();
        }
    }


    //USADA PARA DESERIALIZAR EL JSON QUE ARROJA EL PROCESO

    public class AttributedNodeSyntaxTree
    {
        public string ProductionName { get; set; }
        public Token Token { get; set; }
        public AttributedNodeSyntaxTree[] Childs { get; set; }
        public Dictionary<string,string> Attributes { get; set; }
    }

    //USADA PARA POPULAR LOS CAMPOS DEL TREEVIEW CORRESPONDIENTE
    public class AttributedNodeSyntaxTreeView : TreeNode
    {
        public string ProductionName { get; set; }
        public Token Token { get; set; }
        public string AttributesAsString { get; set; }
        public override string ToString()
        {
            if (Token.Type == 0)
            {
                return "Production Name: " + ProductionName.ToString() + ". \n" + "Atributos: " + AttributesAsString;
            }

            return "Token: \n" + Token.ToString() + "Atributos: " + AttributesAsString;
        }
    }

}


