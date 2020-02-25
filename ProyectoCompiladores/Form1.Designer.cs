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
            this.SuspendLayout();
            // 
            // codeRichTextBox
            // 
            this.codeRichTextBox.AcceptsTab = true;
            this.codeRichTextBox.BackColor = System.Drawing.Color.FromArgb(((int)(((byte)(30)))), ((int)(((byte)(30)))), ((int)(((byte)(30)))));
            this.codeRichTextBox.ForeColor = System.Drawing.Color.White;
            this.codeRichTextBox.Location = new System.Drawing.Point(115, 95);
            this.codeRichTextBox.Name = "codeRichTextBox";
            this.codeRichTextBox.Size = new System.Drawing.Size(575, 303);
            this.codeRichTextBox.TabIndex = 0;
            this.codeRichTextBox.Text = "";
            this.codeRichTextBox.TextChanged += new System.EventHandler(this.codeRichTextBox_TextChanged);
            // 
            // openFileButton
            // 
            this.openFileButton.Location = new System.Drawing.Point(312, 404);
            this.openFileButton.Name = "openFileButton";
            this.openFileButton.Size = new System.Drawing.Size(75, 23);
            this.openFileButton.TabIndex = 1;
            this.openFileButton.Text = "Abrir";
            this.openFileButton.UseVisualStyleBackColor = true;
            this.openFileButton.Click += new System.EventHandler(this.openFileButton_Click);
            // 
            // saveFileAsButton
            // 
            this.saveFileAsButton.Location = new System.Drawing.Point(488, 404);
            this.saveFileAsButton.Name = "saveFileAsButton";
            this.saveFileAsButton.Size = new System.Drawing.Size(111, 23);
            this.saveFileAsButton.TabIndex = 2;
            this.saveFileAsButton.Text = "Guardar como";
            this.saveFileAsButton.UseVisualStyleBackColor = true;
            this.saveFileAsButton.Click += new System.EventHandler(this.saveFileAsButton_Click);
            // 
            // saveFileButton
            // 
            this.saveFileButton.Location = new System.Drawing.Point(393, 404);
            this.saveFileButton.Name = "saveFileButton";
            this.saveFileButton.Size = new System.Drawing.Size(75, 23);
            this.saveFileButton.TabIndex = 3;
            this.saveFileButton.Text = "Guardar";
            this.saveFileButton.UseVisualStyleBackColor = true;
            this.saveFileButton.Click += new System.EventHandler(this.saveFileButton_Click);
            // 
            // fileNameLabel
            // 
            this.fileNameLabel.AutoSize = true;
            this.fileNameLabel.Location = new System.Drawing.Point(124, 35);
            this.fileNameLabel.Name = "fileNameLabel";
            this.fileNameLabel.Size = new System.Drawing.Size(0, 17);
            this.fileNameLabel.TabIndex = 4;
            // 
            // compileButton
            // 
            this.compileButton.Location = new System.Drawing.Point(615, 404);
            this.compileButton.Name = "compileButton";
            this.compileButton.Size = new System.Drawing.Size(75, 23);
            this.compileButton.TabIndex = 5;
            this.compileButton.Text = "Compilar";
            this.compileButton.UseVisualStyleBackColor = true;
            this.compileButton.Click += new System.EventHandler(this.compileButton_Click);
            // 
            // newFileButton
            // 
            this.newFileButton.Location = new System.Drawing.Point(143, 405);
            this.newFileButton.Name = "newFileButton";
            this.newFileButton.Size = new System.Drawing.Size(120, 23);
            this.newFileButton.TabIndex = 6;
            this.newFileButton.Text = "Nuevo archivo";
            this.newFileButton.UseVisualStyleBackColor = true;
            this.newFileButton.Click += new System.EventHandler(this.newFileButton_Click);
            // 
            // Form1
            // 
            this.AutoScaleDimensions = new System.Drawing.SizeF(8F, 16F);
            this.AutoScaleMode = System.Windows.Forms.AutoScaleMode.Font;
            this.ClientSize = new System.Drawing.Size(800, 450);
            this.Controls.Add(this.newFileButton);
            this.Controls.Add(this.compileButton);
            this.Controls.Add(this.fileNameLabel);
            this.Controls.Add(this.saveFileButton);
            this.Controls.Add(this.saveFileAsButton);
            this.Controls.Add(this.openFileButton);
            this.Controls.Add(this.codeRichTextBox);
            this.Name = "Form1";
            this.Text = "Form1";
            this.Load += new System.EventHandler(this.Form1_Load);
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
    }
}

