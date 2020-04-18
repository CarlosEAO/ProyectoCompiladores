namespace ProyectoCompiladores
{
    partial class Form1
    {
        /// <summary>
        /// Variable del diseñador necesaria.
        /// </summary>
        private System.ComponentModel.IContainer components = null;

        /// <summary>
        /// Limpiar los recursos que se estén usando.
        /// </summary>
        /// <param name="disposing">true si los recursos administrados se deben desechar; false en caso contrario.</param>
        protected override void Dispose(bool disposing)
        {
            if (disposing && (components != null))
            {
                components.Dispose();
            }
            base.Dispose(disposing);
        }

        #region Código generado por el Diseñador de Windows Forms

        /// <summary>
        /// Método necesario para admitir el Diseñador. No se puede modificar
        /// el contenido de este método con el editor de código.
        /// </summary>
        private void InitializeComponent()
        {
            this.codeRichTextBox = new System.Windows.Forms.RichTextBox();
            this.openFileButton = new System.Windows.Forms.Button();
            this.saveFileAsButton = new System.Windows.Forms.Button();
            this.saveFileButton = new System.Windows.Forms.Button();
            this.fileNameLabel = new System.Windows.Forms.Label();
            this.compileButton = new System.Windows.Forms.Button();
            this.newFileButton = new System.Windows.Forms.Button();
            this.label2 = new System.Windows.Forms.Label();
            this.tabControl1 = new System.Windows.Forms.TabControl();
            this.tabPage1 = new System.Windows.Forms.TabPage();
            this.tabPage2 = new System.Windows.Forms.TabPage();
            this.tabPage3 = new System.Windows.Forms.TabPage();
            this.tabPage4 = new System.Windows.Forms.TabPage();
            this.tabControl2 = new System.Windows.Forms.TabControl();
            this.tabPage5 = new System.Windows.Forms.TabPage();
            this.tabPage6 = new System.Windows.Forms.TabPage();
            this.menuStrip1 = new System.Windows.Forms.MenuStrip();
            this.tabControl1.SuspendLayout();
            this.tabControl2.SuspendLayout();
            this.SuspendLayout();
            // 
            // codeRichTextBox
            // 
            this.codeRichTextBox.AcceptsTab = true;
            this.codeRichTextBox.BackColor = System.Drawing.Color.FromArgb(((int)(((byte)(64)))), ((int)(((byte)(64)))), ((int)(((byte)(64)))));
            this.codeRichTextBox.ForeColor = System.Drawing.Color.White;
            this.codeRichTextBox.Location = new System.Drawing.Point(17, 66);
            this.codeRichTextBox.Margin = new System.Windows.Forms.Padding(2);
            this.codeRichTextBox.Name = "codeRichTextBox";
            this.codeRichTextBox.Size = new System.Drawing.Size(607, 380);
            this.codeRichTextBox.TabIndex = 0;
            this.codeRichTextBox.Text = "";
            this.codeRichTextBox.TextChanged += new System.EventHandler(this.codeRichTextBox_TextChanged);
            // 
            // openFileButton
            // 
            this.openFileButton.Location = new System.Drawing.Point(111, 20);
            this.openFileButton.Margin = new System.Windows.Forms.Padding(2);
            this.openFileButton.Name = "openFileButton";
            this.openFileButton.Size = new System.Drawing.Size(56, 19);
            this.openFileButton.TabIndex = 1;
            this.openFileButton.Text = "Abrir";
            this.openFileButton.UseVisualStyleBackColor = true;
            this.openFileButton.Click += new System.EventHandler(this.openFileButton_Click);
            // 
            // saveFileAsButton
            // 
            this.saveFileAsButton.Location = new System.Drawing.Point(231, 20);
            this.saveFileAsButton.Margin = new System.Windows.Forms.Padding(2);
            this.saveFileAsButton.Name = "saveFileAsButton";
            this.saveFileAsButton.Size = new System.Drawing.Size(83, 19);
            this.saveFileAsButton.TabIndex = 2;
            this.saveFileAsButton.Text = "Guardar como";
            this.saveFileAsButton.UseVisualStyleBackColor = true;
            this.saveFileAsButton.Click += new System.EventHandler(this.saveFileAsButton_Click);
            // 
            // saveFileButton
            // 
            this.saveFileButton.Location = new System.Drawing.Point(171, 20);
            this.saveFileButton.Margin = new System.Windows.Forms.Padding(2);
            this.saveFileButton.Name = "saveFileButton";
            this.saveFileButton.Size = new System.Drawing.Size(56, 19);
            this.saveFileButton.TabIndex = 3;
            this.saveFileButton.Text = "Guardar";
            this.saveFileButton.UseVisualStyleBackColor = true;
            this.saveFileButton.Click += new System.EventHandler(this.saveFileButton_Click);
            // 
            // fileNameLabel
            // 
            this.fileNameLabel.AutoSize = true;
            this.fileNameLabel.Location = new System.Drawing.Point(380, 23);
            this.fileNameLabel.Margin = new System.Windows.Forms.Padding(2, 0, 2, 0);
            this.fileNameLabel.Name = "fileNameLabel";
            this.fileNameLabel.Size = new System.Drawing.Size(0, 13);
            this.fileNameLabel.TabIndex = 4;
            // 
            // compileButton
            // 
            this.compileButton.Location = new System.Drawing.Point(318, 20);
            this.compileButton.Margin = new System.Windows.Forms.Padding(2);
            this.compileButton.Name = "compileButton";
            this.compileButton.Size = new System.Drawing.Size(56, 19);
            this.compileButton.TabIndex = 5;
            this.compileButton.Text = "Compilar";
            this.compileButton.UseVisualStyleBackColor = true;
            this.compileButton.Click += new System.EventHandler(this.compileButton_Click);
            // 
            // newFileButton
            // 
            this.newFileButton.Location = new System.Drawing.Point(17, 20);
            this.newFileButton.Margin = new System.Windows.Forms.Padding(2);
            this.newFileButton.Name = "newFileButton";
            this.newFileButton.Size = new System.Drawing.Size(90, 19);
            this.newFileButton.TabIndex = 6;
            this.newFileButton.Text = "Nuevo archivo";
            this.newFileButton.UseVisualStyleBackColor = true;
            this.newFileButton.Click += new System.EventHandler(this.newFileButton_Click);
            // 
            // label2
            // 
            this.label2.AutoSize = true;
            this.label2.Location = new System.Drawing.Point(14, 51);
            this.label2.Name = "label2";
            this.label2.Size = new System.Drawing.Size(91, 13);
            this.label2.TabIndex = 10;
            this.label2.Text = "Código a compilar";
            // 
            // tabControl1
            // 
            this.tabControl1.Controls.Add(this.tabPage1);
            this.tabControl1.Controls.Add(this.tabPage2);
            this.tabControl1.Controls.Add(this.tabPage3);
            this.tabControl1.Controls.Add(this.tabPage4);
            this.tabControl1.Location = new System.Drawing.Point(629, 66);
            this.tabControl1.Name = "tabControl1";
            this.tabControl1.SelectedIndex = 0;
            this.tabControl1.Size = new System.Drawing.Size(318, 379);
            this.tabControl1.TabIndex = 11;
            // 
            // tabPage1
            // 
            this.tabPage1.Location = new System.Drawing.Point(4, 22);
            this.tabPage1.Name = "tabPage1";
            this.tabPage1.Padding = new System.Windows.Forms.Padding(3);
            this.tabPage1.Size = new System.Drawing.Size(310, 353);
            this.tabPage1.TabIndex = 0;
            this.tabPage1.Text = "Léxico";
            this.tabPage1.UseVisualStyleBackColor = true;
            // 
            // tabPage2
            // 
            this.tabPage2.Location = new System.Drawing.Point(4, 22);
            this.tabPage2.Name = "tabPage2";
            this.tabPage2.Padding = new System.Windows.Forms.Padding(3);
            this.tabPage2.Size = new System.Drawing.Size(310, 353);
            this.tabPage2.TabIndex = 1;
            this.tabPage2.Text = "Sintáctico";
            this.tabPage2.UseVisualStyleBackColor = true;
            // 
            // tabPage3
            // 
            this.tabPage3.Location = new System.Drawing.Point(4, 22);
            this.tabPage3.Name = "tabPage3";
            this.tabPage3.Size = new System.Drawing.Size(310, 353);
            this.tabPage3.TabIndex = 2;
            this.tabPage3.Text = "Semántico";
            this.tabPage3.UseVisualStyleBackColor = true;
            // 
            // tabPage4
            // 
            this.tabPage4.Location = new System.Drawing.Point(4, 22);
            this.tabPage4.Name = "tabPage4";
            this.tabPage4.Size = new System.Drawing.Size(310, 353);
            this.tabPage4.TabIndex = 3;
            this.tabPage4.Text = "Código intermedio";
            this.tabPage4.UseVisualStyleBackColor = true;
            // 
            // tabControl2
            // 
            this.tabControl2.Controls.Add(this.tabPage5);
            this.tabControl2.Controls.Add(this.tabPage6);
            this.tabControl2.Location = new System.Drawing.Point(17, 451);
            this.tabControl2.Name = "tabControl2";
            this.tabControl2.SelectedIndex = 0;
            this.tabControl2.Size = new System.Drawing.Size(926, 143);
            this.tabControl2.TabIndex = 12;
            // 
            // tabPage5
            // 
            this.tabPage5.Location = new System.Drawing.Point(4, 22);
            this.tabPage5.Name = "tabPage5";
            this.tabPage5.Padding = new System.Windows.Forms.Padding(3);
            this.tabPage5.Size = new System.Drawing.Size(918, 117);
            this.tabPage5.TabIndex = 0;
            this.tabPage5.Text = "Errores";
            this.tabPage5.UseVisualStyleBackColor = true;
            // 
            // tabPage6
            // 
            this.tabPage6.Location = new System.Drawing.Point(4, 22);
            this.tabPage6.Name = "tabPage6";
            this.tabPage6.Padding = new System.Windows.Forms.Padding(3);
            this.tabPage6.Size = new System.Drawing.Size(918, 117);
            this.tabPage6.TabIndex = 1;
            this.tabPage6.Text = "Resultados";
            this.tabPage6.UseVisualStyleBackColor = true;
            // 
            // menuStrip1
            // 
            this.menuStrip1.Location = new System.Drawing.Point(0, 0);
            this.menuStrip1.Name = "menuStrip1";
            this.menuStrip1.Size = new System.Drawing.Size(959, 24);
            this.menuStrip1.TabIndex = 13;
            this.menuStrip1.Text = "menuStrip1";
            // 
            // Form1
            // 
            this.AutoScaleDimensions = new System.Drawing.SizeF(6F, 13F);
            this.AutoScaleMode = System.Windows.Forms.AutoScaleMode.Font;
            this.ClientSize = new System.Drawing.Size(959, 606);
            this.Controls.Add(this.tabControl2);
            this.Controls.Add(this.tabControl1);
            this.Controls.Add(this.label2);
            this.Controls.Add(this.newFileButton);
            this.Controls.Add(this.compileButton);
            this.Controls.Add(this.fileNameLabel);
            this.Controls.Add(this.saveFileButton);
            this.Controls.Add(this.saveFileAsButton);
            this.Controls.Add(this.openFileButton);
            this.Controls.Add(this.codeRichTextBox);
            this.Controls.Add(this.menuStrip1);
            this.MainMenuStrip = this.menuStrip1;
            this.Margin = new System.Windows.Forms.Padding(2);
            this.Name = "Form1";
            this.Text = "Compilador";
            this.Load += new System.EventHandler(this.Form1_Load);
            this.tabControl1.ResumeLayout(false);
            this.tabControl2.ResumeLayout(false);
            this.ResumeLayout(false);
            this.PerformLayout();

        }

        #endregion

        private System.Windows.Forms.RichTextBox codeRichTextBox;
        private System.Windows.Forms.Button openFileButton;
        private System.Windows.Forms.Button saveFileAsButton;
        private System.Windows.Forms.Button saveFileButton;
        private System.Windows.Forms.Label fileNameLabel;
        private System.Windows.Forms.Button compileButton;
        private System.Windows.Forms.Button newFileButton;
        private System.Windows.Forms.Label label2;
        private System.Windows.Forms.TabControl tabControl1;
        private System.Windows.Forms.TabPage tabPage1;
        private System.Windows.Forms.TabPage tabPage2;
        private System.Windows.Forms.TabPage tabPage3;
        private System.Windows.Forms.TabPage tabPage4;
        private System.Windows.Forms.TabControl tabControl2;
        private System.Windows.Forms.TabPage tabPage5;
        private System.Windows.Forms.TabPage tabPage6;
        private System.Windows.Forms.MenuStrip menuStrip1;
    }
}

