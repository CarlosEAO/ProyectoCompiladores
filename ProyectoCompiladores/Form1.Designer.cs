﻿namespace ProyectoCompiladores
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
            this.phasesTabControl = new System.Windows.Forms.TabControl();
            this.lexicTabPage = new System.Windows.Forms.TabPage();
            this.lexicTableLayoutPanel = new System.Windows.Forms.TableLayoutPanel();
            this.tabPage2 = new System.Windows.Forms.TabPage();
            this.parseTreeView = new ProyectoCompiladores.customTreeView();
            this.tabPage3 = new System.Windows.Forms.TabPage();
            this.attributedTreeView = new System.Windows.Forms.TreeView();
            this.tabPage4 = new System.Windows.Forms.TabPage();
            this.resultsTabControl = new System.Windows.Forms.TabControl();
            this.errosTabPage = new System.Windows.Forms.TabPage();
            this.errorsRichTextBox = new System.Windows.Forms.RichTextBox();
            this.resultsTabPage = new System.Windows.Forms.TabPage();
            this.intermediateCodeRichTextBox = new System.Windows.Forms.RichTextBox();
            this.phasesTabControl.SuspendLayout();
            this.lexicTabPage.SuspendLayout();
            this.tabPage2.SuspendLayout();
            this.tabPage3.SuspendLayout();
            this.tabPage4.SuspendLayout();
            this.resultsTabControl.SuspendLayout();
            this.errosTabPage.SuspendLayout();
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
            // phasesTabControl
            // 
            this.phasesTabControl.Controls.Add(this.lexicTabPage);
            this.phasesTabControl.Controls.Add(this.tabPage2);
            this.phasesTabControl.Controls.Add(this.tabPage3);
            this.phasesTabControl.Controls.Add(this.tabPage4);
            this.phasesTabControl.Location = new System.Drawing.Point(629, 66);
            this.phasesTabControl.Name = "phasesTabControl";
            this.phasesTabControl.SelectedIndex = 0;
            this.phasesTabControl.Size = new System.Drawing.Size(532, 379);
            this.phasesTabControl.TabIndex = 11;
            // 
            // lexicTabPage
            // 
            this.lexicTabPage.Controls.Add(this.lexicTableLayoutPanel);
            this.lexicTabPage.Location = new System.Drawing.Point(4, 22);
            this.lexicTabPage.Name = "lexicTabPage";
            this.lexicTabPage.Padding = new System.Windows.Forms.Padding(3);
            this.lexicTabPage.Size = new System.Drawing.Size(524, 353);
            this.lexicTabPage.TabIndex = 0;
            this.lexicTabPage.Text = "Léxico";
            this.lexicTabPage.UseVisualStyleBackColor = true;
            // 
            // lexicTableLayoutPanel
            // 
            this.lexicTableLayoutPanel.ColumnCount = 5;
            this.lexicTableLayoutPanel.ColumnStyles.Add(new System.Windows.Forms.ColumnStyle());
            this.lexicTableLayoutPanel.ColumnStyles.Add(new System.Windows.Forms.ColumnStyle());
            this.lexicTableLayoutPanel.ColumnStyles.Add(new System.Windows.Forms.ColumnStyle());
            this.lexicTableLayoutPanel.ColumnStyles.Add(new System.Windows.Forms.ColumnStyle());
            this.lexicTableLayoutPanel.ColumnStyles.Add(new System.Windows.Forms.ColumnStyle());
            this.lexicTableLayoutPanel.Location = new System.Drawing.Point(0, 0);
            this.lexicTableLayoutPanel.Name = "lexicTableLayoutPanel";
            this.lexicTableLayoutPanel.RowCount = 2;
            this.lexicTableLayoutPanel.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Percent, 50F));
            this.lexicTableLayoutPanel.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Percent, 50F));
            this.lexicTableLayoutPanel.Size = new System.Drawing.Size(524, 353);
            this.lexicTableLayoutPanel.TabIndex = 0;
            // 
            // tabPage2
            // 
            this.tabPage2.Controls.Add(this.parseTreeView);
            this.tabPage2.Location = new System.Drawing.Point(4, 22);
            this.tabPage2.Name = "tabPage2";
            this.tabPage2.Padding = new System.Windows.Forms.Padding(3);
            this.tabPage2.Size = new System.Drawing.Size(524, 353);
            this.tabPage2.TabIndex = 1;
            this.tabPage2.Text = "Sintáctico";
            this.tabPage2.UseVisualStyleBackColor = true;
            // 
            // parseTreeView
            // 
            this.parseTreeView.Location = new System.Drawing.Point(0, 0);
            this.parseTreeView.Margin = new System.Windows.Forms.Padding(2);
            this.parseTreeView.Name = "parseTreeView";
            this.parseTreeView.Size = new System.Drawing.Size(524, 356);
            this.parseTreeView.TabIndex = 0;
            // 
            // tabPage3
            // 
            this.tabPage3.Controls.Add(this.attributedTreeView);
            this.tabPage3.Location = new System.Drawing.Point(4, 22);
            this.tabPage3.Name = "tabPage3";
            this.tabPage3.Size = new System.Drawing.Size(524, 353);
            this.tabPage3.TabIndex = 2;
            this.tabPage3.Text = "Semántico";
            this.tabPage3.UseVisualStyleBackColor = true;
            // 
            // attributedTreeView
            // 
            this.attributedTreeView.Location = new System.Drawing.Point(0, 0);
            this.attributedTreeView.Name = "attributedTreeView";
            this.attributedTreeView.Size = new System.Drawing.Size(524, 353);
            this.attributedTreeView.TabIndex = 0;
            // 
            // tabPage4
            // 
            this.tabPage4.Controls.Add(this.intermediateCodeRichTextBox);
            this.tabPage4.Location = new System.Drawing.Point(4, 22);
            this.tabPage4.Name = "tabPage4";
            this.tabPage4.Size = new System.Drawing.Size(524, 353);
            this.tabPage4.TabIndex = 3;
            this.tabPage4.Text = "Código intermedio";
            this.tabPage4.UseVisualStyleBackColor = true;
            // 
            // resultsTabControl
            // 
            this.resultsTabControl.Controls.Add(this.errosTabPage);
            this.resultsTabControl.Controls.Add(this.resultsTabPage);
            this.resultsTabControl.Location = new System.Drawing.Point(17, 451);
            this.resultsTabControl.Name = "resultsTabControl";
            this.resultsTabControl.SelectedIndex = 0;
            this.resultsTabControl.Size = new System.Drawing.Size(1140, 143);
            this.resultsTabControl.TabIndex = 12;
            // 
            // errosTabPage
            // 
            this.errosTabPage.Controls.Add(this.errorsRichTextBox);
            this.errosTabPage.Location = new System.Drawing.Point(4, 22);
            this.errosTabPage.Name = "errosTabPage";
            this.errosTabPage.Padding = new System.Windows.Forms.Padding(3);
            this.errosTabPage.Size = new System.Drawing.Size(1132, 117);
            this.errosTabPage.TabIndex = 0;
            this.errosTabPage.Text = "Errores";
            this.errosTabPage.UseVisualStyleBackColor = true;
            // 
            // errorsRichTextBox
            // 
            this.errorsRichTextBox.Location = new System.Drawing.Point(0, 0);
            this.errorsRichTextBox.Name = "errorsRichTextBox";
            this.errorsRichTextBox.Size = new System.Drawing.Size(1132, 117);
            this.errorsRichTextBox.TabIndex = 2;
            this.errorsRichTextBox.Text = "";
            // 
            // resultsTabPage
            // 
            this.resultsTabPage.Location = new System.Drawing.Point(4, 22);
            this.resultsTabPage.Name = "resultsTabPage";
            this.resultsTabPage.Padding = new System.Windows.Forms.Padding(3);
            this.resultsTabPage.Size = new System.Drawing.Size(1132, 117);
            this.resultsTabPage.TabIndex = 1;
            this.resultsTabPage.Text = "Resultados";
            this.resultsTabPage.UseVisualStyleBackColor = true;
            // 
            // intermediateCodeRichTextBox
            // 
            this.intermediateCodeRichTextBox.Location = new System.Drawing.Point(1, 3);
            this.intermediateCodeRichTextBox.Name = "intermediateCodeRichTextBox";
            this.intermediateCodeRichTextBox.Size = new System.Drawing.Size(520, 350);
            this.intermediateCodeRichTextBox.TabIndex = 0;
            this.intermediateCodeRichTextBox.Text = "";
            // 
            // Form1
            // 
            this.AutoScaleDimensions = new System.Drawing.SizeF(6F, 13F);
            this.AutoScaleMode = System.Windows.Forms.AutoScaleMode.Font;
            this.ClientSize = new System.Drawing.Size(1173, 606);
            this.Controls.Add(this.resultsTabControl);
            this.Controls.Add(this.phasesTabControl);
            this.Controls.Add(this.label2);
            this.Controls.Add(this.newFileButton);
            this.Controls.Add(this.compileButton);
            this.Controls.Add(this.fileNameLabel);
            this.Controls.Add(this.saveFileButton);
            this.Controls.Add(this.saveFileAsButton);
            this.Controls.Add(this.openFileButton);
            this.Controls.Add(this.codeRichTextBox);
            this.Margin = new System.Windows.Forms.Padding(2);
            this.Name = "Form1";
            this.Text = "Compilador";
            this.Load += new System.EventHandler(this.Form1_Load);
            this.phasesTabControl.ResumeLayout(false);
            this.lexicTabPage.ResumeLayout(false);
            this.tabPage2.ResumeLayout(false);
            this.tabPage3.ResumeLayout(false);
            this.tabPage4.ResumeLayout(false);
            this.resultsTabControl.ResumeLayout(false);
            this.errosTabPage.ResumeLayout(false);
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
        private System.Windows.Forms.TabControl phasesTabControl;
        private System.Windows.Forms.TabPage lexicTabPage;
        private System.Windows.Forms.TabPage tabPage2;
        private System.Windows.Forms.TabPage tabPage3;
        private System.Windows.Forms.TabPage tabPage4;
        private System.Windows.Forms.TabControl resultsTabControl;
        private System.Windows.Forms.TabPage errosTabPage;
        private System.Windows.Forms.TabPage resultsTabPage;
        private System.Windows.Forms.RichTextBox errorsRichTextBox;
        private System.Windows.Forms.TableLayoutPanel lexicTableLayoutPanel;
        private customTreeView parseTreeView;
        private System.Windows.Forms.TreeView attributedTreeView;
        private System.Windows.Forms.RichTextBox intermediateCodeRichTextBox;
    }


    public class customTreeView: System.Windows.Forms.TreeView
    {
        
    }
}

