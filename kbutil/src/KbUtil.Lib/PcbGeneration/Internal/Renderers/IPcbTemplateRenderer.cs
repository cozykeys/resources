namespace KbUtil.Lib.PcbGeneration.Internal
{
    internal interface IPcbTemplateRenderer<in T>
    {
        string KeyboardName { get; set; }

        string Render(T templateData);
    }
}
