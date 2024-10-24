export interface EditorProps {
  cfg: string;
  onChange: (v: string) => void;
  className?: string;
}

export interface AceEditorWrapperProps extends EditorProps {
  maxLines?: number;
  fontSize?: number;
  theme?: Theme;
}

export const themes = [
  "monokai",
  "github",
  "tomorrow",
  "tomorrow_night",
  "kuroir",
  "twilight",
  "xcode",
  "textmate",
  "solarized_dark",
  "solarized_light",
  "terminal",
];
export type Theme = (typeof themes)[number];
