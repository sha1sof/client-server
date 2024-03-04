using System.Net.Sockets;
using System.Text;

namespace client
{
    public partial class Form1 : Form
    {
        private TcpClient client;
        private NetworkStream stream;
        private byte[] buffer = new byte[1024];

        public Form1()
        {
            InitializeComponent();
        }

        private void btnConnect_Click(object sender, EventArgs e)
        {
            try
            {
                string serverAddress = "127.0.0.1";
                int serverPort = 8080;

                client = new TcpClient(serverAddress, serverPort);
                stream = client.GetStream();

                string login = txtLogin.Text;
                byte[] loginData = Encoding.UTF8.GetBytes(login);
                stream.Write(loginData, 0, loginData.Length);

                Thread receiveThread = new Thread(ReceiveMessages);
                receiveThread.Start();

                btnSend.Enabled = true;
                btnConnect.Enabled = false;
            }
            catch (Exception ex)
            {
                MessageBox.Show($"Ошибка при подключении к серверу: {ex.Message}", "Ошибка", MessageBoxButtons.OK, MessageBoxIcon.Error);
            }
        }

        private void ReceiveMessages()
        {
            while (true)
            {
                try
                {
                    int bytesRead = stream.Read(buffer, 0, buffer.Length);
                    if (bytesRead == 0)
                    {
                        MessageBox.Show("Соединение закрыто сервером.", "Информация", MessageBoxButtons.OK, MessageBoxIcon.Information);
                        break;
                    }

                    string message = Encoding.UTF8.GetString(buffer, 0, bytesRead);

                    Invoke(new Action(() => txtInfo.AppendText(message + Environment.NewLine)));
                }
                catch (Exception ex)
                {
                    MessageBox.Show($"Ошибка при получении сообщения: {ex.Message}", "Ошибка", MessageBoxButtons.OK, MessageBoxIcon.Error);
                    break;
                }
            }
        }


        private void btnSend_Click(object sender, EventArgs e)
        {
            if (client != null && client.Connected)
            {
                try
                {
                    string receiverLogin = txtReceiverId.Text;
                    string message = txtMessage.Text;

                    string fullMessage = receiverLogin + " " + message;
                    byte[] data = Encoding.UTF8.GetBytes(fullMessage);

                    stream.Write(data, 0, data.Length);

                    txtInfo.Text += $"Me to {receiverLogin}: {message}{Environment.NewLine}";
                    txtMessage.Text = string.Empty;
                }
                catch (Exception ex)
                {
                    MessageBox.Show($"Ошибка при отправке сообщения: {ex.Message}", "Ошибка", MessageBoxButtons.OK, MessageBoxIcon.Error);
                }
            }
        }

        private void Form1_FormClosing(object sender, FormClosingEventArgs e)
        {
            if (client != null)
            {
                client.Close();
            }
        }
    }
}