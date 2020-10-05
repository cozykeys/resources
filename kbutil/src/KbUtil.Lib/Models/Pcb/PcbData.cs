namespace KbUtil.Lib.Models.Pcb
{
    using System.Collections.Generic;

    internal class PcbData
    {
        public PcbData(IEnumerable<Switch> switches)
        {
            InitializeSwitchMatrix(switches);
            InitializeNetDictionary();
        }

        public Switch[][] Switches { get; private set; }
        public int RowCount { get; private set; }
        public int ColumnCount { get; private set; }
        public Dictionary<string, int> NetDictionary { get; private set; }

        private void InitializeSwitchMatrix(IEnumerable<Switch> switches)
        {
            int maxRowIndex = int.MinValue;
            int maxColumnIndex = int.MinValue;

            foreach (var @switch in switches)
            {
                if (@switch.Row > maxRowIndex)
                {
                    maxRowIndex = @switch.Row;
                }

                if (@switch.Column > maxColumnIndex)
                {
                    maxColumnIndex = @switch.Column;
                }
            }

            RowCount = maxRowIndex + 1;
            ColumnCount = maxColumnIndex + 1;

            Switches = new Switch[RowCount][];

            for (int i = 0; i < RowCount; ++i)
            {
                Switches[i] = new Switch[ColumnCount];
            }

            for (int i = 0; i < RowCount; ++i)
            {
                for (int j = 0; j < RowCount; ++j)
                {
                    Switches[i][j] = null;
                }
            }

            foreach (var @switch in switches)
            {
                Switches[@switch.Row][@switch.Column] = @switch;
            }
        }

        private void InitializeNetDictionary()
        {
            int index = 0;
            NetDictionary = new Dictionary<string, int>
            {
                { "\"\"", index++ }
            };

            NetDictionary["N-5V-0"] = index++;
            NetDictionary["N-GND-0"] = index++;
            NetDictionary["N-MOSFET-0"] = index++;

            for (int i = 0; i < RowCount; ++i)
            {
                NetDictionary[$"N-row-{i}"] = index++;
            }

            for (int i = 0; i < ColumnCount; ++i)
            {
                NetDictionary[$"N-col-{i}"] = index++;
            }

            for (int i = 0; i < RowCount; ++i)
            {
                for (int j = 0; j < ColumnCount; ++j)
                {
                    if (Switches[i][j] != null)
                    {
                        NetDictionary[$"N-diode-{i}-{j}"] = index++;
                    }
                }
            }

            for (int i = 0; i < RowCount; ++i)
            {
                for (int j = 0; j < ColumnCount; ++j)
                {
                    if (Switches[i][j] != null)
                    {
                        NetDictionary[$"N-LED-{i}-{j}"] = index++;
                    }
                }
            }

            for (int i = 0; i < 12; ++i)
            {
                NetDictionary[$"N-RGB-D{i}"] = index++;
            }

            NetDictionary["N-LED-PWM"] = index++;
        }
    }
}
