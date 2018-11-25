namespace KbUtil.Lib.PcbGeneration.Internal
{
    internal interface IPcbTemplateRenderer<in T>
    {
        string Render(T templateData);
    }
}
