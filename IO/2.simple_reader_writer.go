package IO

// SimpleWriterSolution Read ---> Writer(单) 解决方案
type SimpleWriterSolution interface {
	Copy(dst Writer, src Reader) (written int64, err error) // io也提供了io.Copy
}

// reader -----> writer

/* 直接通过 io.Copy 方式将responseReader中的数据给到 Writer   reader -----> Writer
func handleGetInfoRefsWithGitAly(ctx context.Context, responseWriter *HttpResponseWriter, a *api.Response, rpc, gitProtocol, encoding string) error {
        ...
        infoRefsResponseReader, err := smartHttp.InfoRefsResponseReader(ctx, &a.Repository, rpc, gitConfigOptions(a), gitProtocol)
        ...
        if _, err = io.Copy(w, infoRefsResponseReader); err != nil {
            return err
        }
        ...
}
*/
