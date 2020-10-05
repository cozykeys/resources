namespace KbUtil.Lib.PcbGeneration.Internal.Extensions
{
    using System.Text.RegularExpressions;

    internal static class StringExtensions
    {
        public static string StripComments(this string raw)
            => Regex.Replace(raw, @"/\*.*\*/", "");
    }
}
