namespace client
{
    partial class Form1
    {
        /// <summary>
        ///  Required designer variable.
        /// </summary>
        private System.ComponentModel.IContainer components = null;

        /// <summary>
        ///  Clean up any resources being used.
        /// </summary>
        /// <param name="disposing">true if managed resources should be disposed; otherwise, false.</param>
        protected override void Dispose(bool disposing)
        {
            if (disposing && (components != null))
            {
                components.Dispose();
            }
            base.Dispose(disposing);
        }

        #region Windows Form Designer generated code

        /// <summary>
        ///  Required method for Designer support - do not modify
        ///  the contents of this method with the code editor.
        /// </summary>
        private void InitializeComponent()
        {
            label1 = new Label();
            btnConnect = new Button();
            txtInfo = new TextBox();
            txtMessage = new TextBox();
            label2 = new Label();
            btnSend = new Button();
            txtReceiverId = new TextBox();
            txtLogin = new TextBox();
            checkBoxEncryption = new CheckBox();
            label3 = new Label();
            SuspendLayout();
            // 
            // label1
            // 
            label1.AutoSize = true;
            label1.Font = new Font("Segoe UI", 20.25F, FontStyle.Regular, GraphicsUnit.Point, 204);
            label1.Location = new Point(7, 9);
            label1.Name = "label1";
            label1.Size = new Size(195, 37);
            label1.TabIndex = 0;
            label1.Text = "Введите логин";
            // 
            // btnConnect
            // 
            btnConnect.Anchor = AnchorStyles.Top | AnchorStyles.Right;
            btnConnect.Location = new Point(598, 23);
            btnConnect.Name = "btnConnect";
            btnConnect.Size = new Size(75, 23);
            btnConnect.TabIndex = 2;
            btnConnect.Text = "Connect";
            btnConnect.UseVisualStyleBackColor = true;
            btnConnect.Click += btnConnect_Click;
            // 
            // txtInfo
            // 
            txtInfo.Anchor = AnchorStyles.Top | AnchorStyles.Bottom | AnchorStyles.Left | AnchorStyles.Right;
            txtInfo.Location = new Point(118, 52);
            txtInfo.Multiline = true;
            txtInfo.Name = "txtInfo";
            txtInfo.ScrollBars = ScrollBars.Both;
            txtInfo.Size = new Size(540, 276);
            txtInfo.TabIndex = 3;
            // 
            // txtMessage
            // 
            txtMessage.Anchor = AnchorStyles.Bottom | AnchorStyles.Left;
            txtMessage.Location = new Point(118, 362);
            txtMessage.Name = "txtMessage";
            txtMessage.Size = new Size(540, 23);
            txtMessage.TabIndex = 5;
            // 
            // label2
            // 
            label2.Anchor = AnchorStyles.Bottom | AnchorStyles.Left;
            label2.AutoSize = true;
            label2.Font = new Font("Segoe UI", 14.25F, FontStyle.Regular, GraphicsUnit.Point, 204);
            label2.Location = new Point(7, 360);
            label2.Name = "label2";
            label2.Size = new Size(115, 25);
            label2.TabIndex = 4;
            label2.Text = "Сообщение";
            // 
            // btnSend
            // 
            btnSend.Anchor = AnchorStyles.Bottom | AnchorStyles.Right;
            btnSend.Location = new Point(583, 391);
            btnSend.Name = "btnSend";
            btnSend.Size = new Size(75, 23);
            btnSend.TabIndex = 6;
            btnSend.Text = "Написать";
            btnSend.UseVisualStyleBackColor = true;
            btnSend.Click += btnSend_Click;
            // 
            // txtReceiverId
            // 
            txtReceiverId.Anchor = AnchorStyles.Bottom | AnchorStyles.Left;
            txtReceiverId.Location = new Point(118, 334);
            txtReceiverId.Name = "txtReceiverId";
            txtReceiverId.Size = new Size(100, 23);
            txtReceiverId.TabIndex = 7;
            // 
            // txtLogin
            // 
            txtLogin.Anchor = AnchorStyles.Top | AnchorStyles.Left | AnchorStyles.Right;
            txtLogin.Location = new Point(208, 23);
            txtLogin.Name = "txtLogin";
            txtLogin.Size = new Size(368, 23);
            txtLogin.TabIndex = 8;
            // 
            // checkBoxEncryption
            // 
            checkBoxEncryption.Anchor = AnchorStyles.Top | AnchorStyles.Bottom | AnchorStyles.Left;
            checkBoxEncryption.AutoSize = true;
            checkBoxEncryption.Location = new Point(12, 94);
            checkBoxEncryption.Name = "checkBoxEncryption";
            checkBoxEncryption.Size = new Size(90, 19);
            checkBoxEncryption.TabIndex = 9;
            checkBoxEncryption.Text = "Шифровать";
            checkBoxEncryption.UseVisualStyleBackColor = true;
            // 
            // label3
            // 
            label3.Anchor = AnchorStyles.Bottom | AnchorStyles.Left;
            label3.AutoSize = true;
            label3.Font = new Font("Segoe UI", 14.25F, FontStyle.Regular, GraphicsUnit.Point, 204);
            label3.Location = new Point(56, 331);
            label3.Name = "label3";
            label3.Size = new Size(56, 25);
            label3.TabIndex = 11;
            label3.Text = "Кому";
            // 
            // Form1
            // 
            AutoScaleDimensions = new SizeF(7F, 15F);
            AutoScaleMode = AutoScaleMode.Font;
            ClientSize = new Size(679, 422);
            Controls.Add(label3);
            Controls.Add(checkBoxEncryption);
            Controls.Add(txtLogin);
            Controls.Add(txtReceiverId);
            Controls.Add(btnSend);
            Controls.Add(txtMessage);
            Controls.Add(label2);
            Controls.Add(txtInfo);
            Controls.Add(btnConnect);
            Controls.Add(label1);
            MaximizeBox = false;
            Name = "Form1";
            StartPosition = FormStartPosition.CenterScreen;
            Text = "TCP client";
            ResumeLayout(false);
            PerformLayout();
        }

        #endregion

        private Label label1;
        private Button btnConnect;
        private TextBox txtInfo;
        private TextBox txtMessage;
        private Label label2;
        private Button btnSend;
        private TextBox txtReceiverId;
        private TextBox txtLogin;
        private CheckBox checkBoxEncryption;
        private Label label3;
    }
}
