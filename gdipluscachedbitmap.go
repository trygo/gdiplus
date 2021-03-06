package gdiplus

type CachedBitmap struct {
	GdiplusBase
	nativeCachedBitmap *GpCachedBitmap
}

func NewCachedBitmap(bitmap *Bitmap, graphics *Graphics) (*CachedBitmap, error) {
	cachedBitmap := &CachedBitmap{}
	cachedBitmap.setStatus(GdipCreateCachedBitmap((*GpBitmap)(bitmap.nativeImage), graphics.nativeGraphics, &cachedBitmap.nativeCachedBitmap))
	return cachedBitmap, cachedBitmap.LastError
}

func (this *CachedBitmap) Release() {
	if this.nativeCachedBitmap != nil {
		this.setStatus(GdipDeleteCachedBitmap(this.nativeCachedBitmap))
		this.nativeCachedBitmap = nil
	}
}
